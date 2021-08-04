package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {

		app.Use(fibertrace.Middleware(
			fibertrace.WithServiceName(datadog.Service()),
			fibertrace.WithAnalyticsRate(datadog.AnalyticsRate()),
		))

		logger.Debug("datadog middleware successfully enabled in fiber")

		return nil

	}
}
