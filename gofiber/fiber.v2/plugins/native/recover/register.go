package recover

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling recover middleware in fiber")

	app.Use(recover.New())

	logger.Debug("recover middleware successfully enabled in fiber")

	return nil
}
