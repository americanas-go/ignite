package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v7"
)

type Datadog struct {
	options *Options
}

func NewDatadogWithConfigPath(path string) (*Datadog, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func (d *Datadog) Register(ctx context.Context, client *redis.Client) error {

	if !d.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")

	redistrace.WrapClient(client)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}

func Register(ctx context.Context, client *redis.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	d := NewDatadogWithOptions(o)
	return d.Register(ctx, client)
}
