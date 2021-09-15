package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

type Health struct {
	options *Options
}

func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}

	plugin := NewHealthWithOptions(o)
	return plugin.Register(ctx, client)
}

func (i *Health) Register(ctx context.Context, client *resty.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating resty in health")

	checker := NewChecker(client, i.options)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("resty successfully integrated in health")

	return nil
}
