package opentracing

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewOpentracingWithOptions(o)
	return n.Register(ctx, options)
}

type Opentracing struct {
	options *Options
}

func NewOpentracingWithConfigPath(path string) (*Opentracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpentracingWithOptions(o), nil
}

func NewOpentracingWithOptions(options *Options) *Opentracing {
	return &Opentracing{options: options}
}

func (d *Opentracing) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling opentracing middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(opentracingMiddleware())

		logger.Debug("opentracing middleware successfully enabled in fiber")

		return nil
	}
}
