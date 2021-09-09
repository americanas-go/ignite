package aws

import (
	"context"

	"github.com/americanas-go/ignite/net/http/client"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	c "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// Plugin defines a function to process plugin.
type Plugin func(context.Context, *aws.Config) error

// NewConfig returns aws config.
func NewConfig(ctx context.Context, plugins ...Plugin) aws.Config {

	o, err := NewOptions()
	if err != nil {
		panic(err)
	}

	return NewConfigWithOptions(ctx, o, plugins...)
}

// NewConfigWithOptions returns aws config with options.
func NewConfigWithOptions(ctx context.Context, options *Options, plugins ...Plugin) aws.Config {

	logger := log.FromContext(ctx)

	cfg, err := c.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Panicf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}
	}

	if options.DefaultRegion != "" {
		cfg.Region = options.DefaultRegion
	}

	if options.AccessKeyId != "" && options.SecretAccessKey != "" {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(options.AccessKeyId, options.SecretAccessKey, options.SessionToken)
	}

	httpClient := client.NewClientWithOptions(ctx, &options.HttpClient)

	cfg.Retryer = retryerConfig(options)
	cfg.HTTPClient = httpClient

	for _, plugin := range plugins {
		if err := plugin(ctx, &cfg); err != nil {
			panic(err)
		}
	}

	return cfg
}

func retryerConfig(options *Options) func() aws.Retryer {

	return func() aws.Retryer {

		return retry.NewStandard(func(o *retry.StandardOptions) {

			o.MaxAttempts = options.MaxAttempts

			if !options.HasRateLimit {
				o.RateLimiter = noRateLimit{}
			}

		})
	}
}

type noRateLimit struct{}

func (noRateLimit) AddTokens(uint) error { return nil }

func (noRateLimit) GetToken(context.Context, uint) (func() error, error) { return nil, nil }
