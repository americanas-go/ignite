package requestid

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling requestID middleware in fiber")

	app.Use(requestid.New())

	logger.Debug("requestID middleware successfully enabled in fiber")

	return nil
}
