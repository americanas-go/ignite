package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v7"
)

type ClusterDatadog struct {
	options *Options
}

func NewClusterDatadogWithConfigPath(path string, traceOptions ...redistrace.ClientOption) (*ClusterDatadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewClusterDatadogWithOptions(o), nil
}

func NewClusterDatadog(traceOptions ...redistrace.ClientOption) *ClusterDatadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClusterDatadogWithOptions(o)
}

func NewClusterDatadogWithOptions(options *Options) *ClusterDatadog {
	return &ClusterDatadog{options: options}
}

func (d *ClusterDatadog) Register(ctx context.Context, client *redis.ClusterClient) error {

	if !d.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")

	redistrace.WrapClient(client, d.options.TraceOptions...)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}

func ClusterRegister(ctx context.Context, client *redis.ClusterClient) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	d := NewClusterDatadogWithOptions(o)
	return d.Register(ctx, client)
}
