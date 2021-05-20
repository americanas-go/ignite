package status

import (
	"context"
	"fmt"
	"net/http"

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

	statusRoute := getRoute()

	logger.Tracef("configuring status router on %s in fiber", statusRoute)

	return nil, func(ctx context.Context, app *f.App) error {

		app.Get(statusRoute, func(c *f.Ctx) error {

			c = c.Status(http.StatusOK)

			resourceStatus := response.NewResourceStatus()

			if options.Type != "REST" {
				return c.SendString(fmt.Sprintf("%v", resourceStatus))
			}

			return c.JSON(resourceStatus)
		})

		logger.Debugf("status router configured on %s in fiber", statusRoute)
		return nil
	}

}
