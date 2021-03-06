package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

type ClientNewrelic struct {
	options *Options
}

func NewClientNewrelicWithConfigPath(path string) (*ClientNewrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientNewrelicWithOptions(o), nil
}

func NewClientNewrelicWithOptions(options *Options) *ClientNewrelic {
	return &ClientNewrelic{options: options}
}

func (d *ClientNewrelic) Register(ctx context.Context, client *redis.Client) error {

	if !d.options.Enabled || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in newrelic")

	client.AddHook(nrredis.NewHook(client.Options()))

	logger.Debug("redis successfully integrated in newrelic")

	return nil
}

func ClientRegister(ctx context.Context, client *redis.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	n := NewClientNewrelicWithOptions(o)
	return n.Register(ctx, client)
}
