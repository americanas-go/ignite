package pprof

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

// Register registers a new pprof plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewPProfWithOptions(o)
	return n.Register(ctx, options)
}

// PProf represents a new pprof plugin for fiber.
type PProf struct {
	options *Options
}

// NewPProfWithConfigPath returns a new pprof plugin with options from config path.
func NewPProfWithConfigPath(path string) (*PProf, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPProfWithOptions(o), nil
}

// NewPProfWithOptions returns a new pprof plugin with options.
func NewPProfWithOptions(options *Options) *PProf {
	return &PProf{options: options}
}

// Register registers this pprof plugin for fiber.
func (d *PProf) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling pprof middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(pprof.New())

		logger.Debug("pprof middleware successfully enabled in fiber")

		return nil
	}
}
