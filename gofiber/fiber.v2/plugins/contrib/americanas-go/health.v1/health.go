package health

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	f "github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	health := NewHealthWithOptions(o)
	return health.Register(ctx, options)
}

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

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
