package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	checkRoute := getRoute()

	logger.Tracef("configuring multi server check router on %s in fiber", checkRoute)

	return nil, func(ctx context.Context, app *f.App) error {

		app.Get(checkRoute, handler)

		logger.Debugf("multi server check router configured on %s in fiber", checkRoute)
		return nil
	}

}

func handler(c *f.Ctx) error {

	status := http.StatusOK
	msg := "OK"

	if err := multiserver.Check(c.Context()); err != nil {
		status = http.StatusServiceUnavailable
		msg = "Service Unavailable"
	}

	return c.Status(status).SendString(msg)
}
