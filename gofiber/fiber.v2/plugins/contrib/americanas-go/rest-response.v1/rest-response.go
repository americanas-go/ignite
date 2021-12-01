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

// Register registers a new status plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	l := NewStatus()
	return l.Register(ctx, options)
}

// Status represents a status plugin for fiber.
type Status struct {
	options *Options
}

// NewStatusWithOptions returns a new status plugin with options.
func NewStatusWithOptions(options *Options) *Status {
	return &Status{options: options}
}

// NewStatusWithConfigPath returns a new status plugin with options from config path.
func NewStatusWithConfigPath(path string) (*Status, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStatusWithOptions(o), nil
}

// NewStatusWithOptions returns a new status plugin with default options.
func NewStatus() *Status {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewStatusWithOptions(o)
}

// Register registers this status plugin for fiber.
func (i *Status) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	statusRoute := i.options.Route

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
