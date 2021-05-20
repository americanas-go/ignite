package requestid

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling requestID middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(requestid.New())

		logger.Debug("requestID middleware successfully enabled in fiber")

		return nil
	}
}
