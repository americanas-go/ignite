package opentracing

import (
	"context"

	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
)

// Register registers a new opentracing plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewOpentracingWithOptions(o)
	return n.Register(ctx, options)
}

// Opentracing represents a new opentracing plugin for fiber.
type Opentracing struct {
	options *Options
}

// NewOpentracingWithConfigPath returns a new opentracing plugin with options from config path.
func NewOpentracingWithConfigPath(path string) (*Opentracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpentracingWithOptions(o), nil
}

// NewOpentracingWithOptions returns a new opentracing plugin with options.
func NewOpentracingWithOptions(options *Options) *Opentracing {
	return &Opentracing{options: options}
}

// Register registers this opentracing plugin for fiber.
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
