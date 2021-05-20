package newrelic

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling newrelic middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(middleware(newrelic.Application()))

		logger.Debug("newrelic middleware successfully enabled in fiber")

		return nil
	}
}
