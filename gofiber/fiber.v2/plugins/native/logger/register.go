package logger

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	l := log.FromContext(ctx)
	l.Trace("enabling logger middleware in fiber")

	app.Use(logger.New(logger.Config{
		Output: log.GetLogger().Output(),
	}))

	l.Debug("logger middleware successfully enabled in fiber")

	return nil
}
