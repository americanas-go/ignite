package status

import (
	"context"
	"net/http"

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

	statusRoute := getRoute()

	logger.Tracef("configuring status router on %s in fiber", statusRoute)

	app.Get(statusRoute, func(c *f.Ctx) error {
		return fiber.JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
	})

	logger.Debugf("status router configured on %s in fiber", statusRoute)

	return nil
}
