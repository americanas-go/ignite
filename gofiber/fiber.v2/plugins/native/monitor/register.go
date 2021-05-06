package monitor

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling monitor middleware in fiber")

	app.Use(monitor.New())

	logger.Debug("monitor middleware successfully enabled in fiber")

	return nil
}
