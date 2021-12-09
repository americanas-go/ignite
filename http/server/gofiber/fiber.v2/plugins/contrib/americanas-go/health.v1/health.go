package health

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	f "github.com/gofiber/fiber/v2"
)

// Register registers a new health checker for fiber with options.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	health := NewHealthWithOptions(o)
	return health.Register(ctx, options)
}

// Health represets a health checker plugin for fiber
type Health struct {
	options *Options
}

// NewHealthWithOptions returns a health checker with options.
func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

// NewHealthWithOptions returns a health checker with options from config path.
func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

// NewHealthWithOptions returns a health checker with options.
func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

// Register registers this health checker for fiber with options.
func (i *Health) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	healthRoute := i.options.Route

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
