package monitor

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewMonitorWithOptions(o)
	return n.Register(ctx, options)
}

type Monitor struct {
	options *Options
}

func NewMonitorWithConfigPath(path string) (*Monitor, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewMonitorWithOptions(o), nil
}

func NewMonitorWithOptions(options *Options) *Monitor {
	return &Monitor{options: options}
}

func (d *Monitor) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling monitor middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(monitor.New())

		logger.Debug("monitor middleware successfully enabled in fiber")

		return nil
	}
}
