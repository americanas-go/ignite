package newrelic

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewNewrelicWithOptions(o)
	return n.Register(ctx, options)
}

type Newrelic struct {
	options *Options
}

func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

func (d *Newrelic) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !d.options.Enabled || !newrelic.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling newrelic middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(middleware(newrelic.Application()))

		logger.Debug("newrelic middleware successfully enabled in fiber")

		return nil
	}
}
