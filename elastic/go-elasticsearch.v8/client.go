package elasticsearch

import (
	"context"
	"strings"
	"time"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// Plugin defines a function to process plugin.
type Plugin func(context.Context, *elasticsearch.Client) error

// NewClient returns elasticsearch client with default options.
func NewClient(ctx context.Context, plugins ...Plugin) (*elasticsearch.Client, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClientWithOptions(ctx, o, plugins...)
}

// NewClientWithConfigPath returns elasticsearch client with options from config path.
func NewClientWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*elasticsearch.Client, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientWithOptions(ctx, opts, plugins...)
}

// NewClientWithOptions returns elasticsearch client with options.
func NewClientWithOptions(ctx context.Context, o *Options, plugins ...Plugin) (client *elasticsearch.Client, err error) {

	logger := log.FromContext(ctx)

	cfg := elasticsearch.Config{
		Addresses:             o.Addresses,
		Username:              o.Username,
		Password:              o.Password,
		CloudID:               o.CloudID,
		APIKey:                o.APIKey,
		RetryOnStatus:         o.RetryOnStatus,
		DisableRetry:          o.DisableRetry,
		EnableRetryOnTimeout:  o.EnableRetryOnTimeout,
		MaxRetries:            o.MaxRetries,
		DiscoverNodesOnStart:  o.DiscoverNodesOnStart,
		DiscoverNodesInterval: o.DiscoverNodesInterval,
		EnableMetrics:         o.EnableMetrics,
		EnableDebugLogger:     o.EnableDebugLogger,
		RetryBackoff:          backOff,
		Logger:                &Logger{},
	}

	if o.CACert != "" {
		cfg.CACert = []byte(o.CACert)
	}

	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	var res *esapi.Response

	res, err = client.Ping(client.Ping.WithPretty())
	if err != nil {
		return nil, err
	}

	for _, plugin := range plugins {
		if err := plugin(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to Elastic Search server: %v status: %s", strings.Join(o.Addresses, ","), res.Status())

	return client, err
}

func backOff(attempt int) time.Duration {
	b := config.Duration(retryBackoff)
	return time.Duration(attempt) * b
}
