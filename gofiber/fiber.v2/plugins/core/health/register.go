package health

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *f.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	healthRoute := getRoute()

	logger.Tracef("configuring health router on %s in fiber", healthRoute)

	app.Get(healthRoute, func(c *f.Ctx) error {

		ctx, cancel := context.WithCancel(c.Context())
		defer cancel()

		resp, httpCode := response.NewHealth(ctx)

		return fiber.JSON(c, httpCode, resp, nil)
	})

	logger.Debugf("health router configured on %s in fiber", healthRoute)

	return nil
}
