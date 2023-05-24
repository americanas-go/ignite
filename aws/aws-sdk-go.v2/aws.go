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
func NewConfig(ctx context.Context, plugins ...Plugin) (aws.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return aws.Config{}, err
	}

	return NewConfigWithOptions(ctx, o, plugins...)
}

// NewConfigWithConfigPath NewConfigWithOptions returns aws config with options from config path.
func NewConfigWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (aws.Config, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return aws.Config{}, err
	}
	return NewConfigWithOptions(ctx, opts, plugins...)
}

// NewConfigWithOptions returns aws config with options.
func NewConfigWithOptions(ctx context.Context, options *Options, plugins ...Plugin) (aws.Config, error) {

	logger := log.FromContext(ctx)

	cfg, err := c.LoadDefaultConfig(ctx, c.WithEndpointResolverWithOptions(customResolver(ctx, options)))
	if err != nil {
		logger.Errorf("unable to load AWS SDK config, %s", err.Error())
		return aws.Config{}, nil
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
			return aws.Config{}, err
		}
	}

	return cfg, nil
}

func customResolver(ctx context.Context, options *Options) aws.EndpointResolverWithOptionsFunc {

	logger := log.FromContext(ctx)

	return func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
		if ce, ok := options.CustomEndpoint[service]; ok {

			logger.Debugf("configuring custom endpoint to service %s", service)

			return aws.Endpoint{
				PartitionID:       ce.PartitionID,
				URL:               ce.URL,
				SigningRegion:     ce.SigningRegion,
				HostnameImmutable: ce.HostnameImmutable,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	}
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
