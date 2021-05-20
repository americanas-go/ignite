package logger

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !IsEnabled() {
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
