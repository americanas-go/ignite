package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v8"
)

type ClientHealth struct {
	options *Options
}

func NewClientHealthWithOptions(options *Options) *ClientHealth {
	return &ClientHealth{options: options}
}

func NewClientHealth() *ClientHealth {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClientHealthWithOptions(o)
}

func (i *ClientHealth) Register(ctx context.Context, client *redis.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClientChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}
