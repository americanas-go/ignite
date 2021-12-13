package prometheus

import (
	"context"

	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	"github.com/ansrivas/fiberprometheus/v2"
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

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling prometheus middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		prometheus := fiberprometheus.New("")
		app.Use(prometheus.Middleware)

		logger.Debug("prometheus middleware successfully enabled in fiber")

		prometheusRoute := d.options.Route

		logger.Tracef("configuring prometheus metric router on %s in fiber", prometheusRoute)

		prometheus.RegisterAt(app, prometheusRoute)

		logger.Debugf("prometheus metric router configured on %s in fiber", prometheusRoute)

		return nil
	}
}
