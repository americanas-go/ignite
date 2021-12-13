package monitor

import (
	"context"

	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Register registers a new monitor plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewMonitorWithOptions(o)
	return n.Register(ctx, options)
}

// Monitor represents a new monitor plugin for fiber.
type Monitor struct {
	options *Options
}

// NewMonitorWithConfigPath returns a new monitor plugin with options from config path.
func NewMonitorWithConfigPath(path string) (*Monitor, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewMonitorWithOptions(o), nil
}

// NewMonitorWithOptions returns a new monitor plugin with options.
func NewMonitorWithOptions(options *Options) *Monitor {
	return &Monitor{options: options}
}

// Register registers this monitor plugin for fiber.
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
