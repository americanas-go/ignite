package logger

import (
	"context"

	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Register registers a new logger plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewLoggerWithOptions(o)
	return n.Register(ctx, options)
}

// Logger represents a new logger plugin for fiber.
type Logger struct {
	options *Options
}

// NewLoggerWithConfigPath returns a new logger plugin with options from config path.
func NewLoggerWithConfigPath(path string) (*Logger, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewLoggerWithOptions(o), nil
}

// NewLoggerWithOptions returns a new logger plugin with options.
func NewLoggerWithOptions(options *Options) *Logger {
	return &Logger{options: options}
}

// Register registers this logger plugin for fiber.
func (d *Logger) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !d.options.Enabled {
		return nil, nil
	}

	l := log.FromContext(ctx)
	l.Trace("enabling logger middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(logger.New(logger.Config{
			Output: log.GetLogger().Output(),
		}))

		l.Debug("logger middleware successfully enabled in fiber")

		return nil
	}
}
