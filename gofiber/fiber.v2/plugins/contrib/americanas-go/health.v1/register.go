package health

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	healthRoute := getRoute()

	logger.Tracef("configuring health router on %s in fiber", healthRoute)

	return nil, func(ctx context.Context, app *f.App) error {

		app.Get(healthRoute, func(c *f.Ctx) error {

			ctx, cancel := context.WithCancel(c.Context())
			defer cancel()

			resp, httpCode := response.NewHealth(ctx)

			c = c.Status(httpCode)

			if options.Type != "REST" {
				return c.SendString(resp.Status.String())
			}

			return c.JSON(resp)
		})

		logger.Debugf("health router configured on %s in fiber", healthRoute)

		return nil
	}
}
