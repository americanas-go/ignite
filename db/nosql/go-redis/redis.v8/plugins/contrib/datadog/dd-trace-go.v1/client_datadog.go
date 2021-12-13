package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/apm/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v8"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
)

// ClientDatadog represents a datadog client for redis.
type ClientDatadog struct {
	options *Options
}

// NewClientDatadogWithConfigPath returns a new datadog client with options from config path.
func NewClientDatadogWithConfigPath(path string, traceOptions ...redistrace.ClientOption) (*ClientDatadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}

	if !datadog.IsTracerEnabled() {
		o.Enabled = false
	}

	return NewClientDatadogWithOptions(o), nil
}

// NewClientDatadog returns a new datadog client with default options.
func NewClientDatadog(traceOptions ...redistrace.ClientOption) (*ClientDatadog, error) {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	return NewClientDatadogWithOptions(o), nil
}

// NewClientDatadogWithOptions returns a new datadog client with options.
func NewClientDatadogWithOptions(options *Options) *ClientDatadog {
	return &ClientDatadog{options: options}
}

// Register registers this datadog client to redis client.
func (d *ClientDatadog) Register(ctx context.Context, client *redis.Client) error {
	if !d.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")

	redistrace.WrapClient(client, d.options.TraceOptions...)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}

// ClientRegister registers a new datadog client to redis client.
func ClientRegister(ctx context.Context, client *redis.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	d := NewClientDatadogWithOptions(o)
	return d.Register(ctx, client)
}
