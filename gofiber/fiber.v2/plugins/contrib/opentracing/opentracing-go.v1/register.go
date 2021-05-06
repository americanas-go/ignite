package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	l := log.FromContext(ctx)
	l.Trace("enabling opentracing middleware in fiber")

	app.Use(opentracingMiddleware())

	l.Debug("opentracing middleware successfully enabled in fiber")

	return nil
}
