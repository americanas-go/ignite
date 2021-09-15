package cors

import (
	"context"
	"strings"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewCorsWithOptions(o)
	return n.Register(ctx, options)
}

type Cors struct {
	options *Options
}

func NewCorsWithConfigPath(path string) (*Cors, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCorsWithOptions(o), nil
}

func NewCorsWithOptions(options *Options) *Cors {
	return &Cors{options: options}
}

func (d *Cors) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling cors middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     strings.Join(d.options.Allowed.Origins, ","),
			AllowMethods:     strings.Join(d.options.Allowed.Methods, ","),
			AllowHeaders:     strings.Join(d.options.Allowed.Headers, ","),
			AllowCredentials: d.options.Allowed.Credentials,
			ExposeHeaders:    strings.Join(d.options.Exposed.Headers, ","),
			MaxAge:           d.options.MaxAge,
		}))

		logger.Debug("cors middleware successfully enabled in fiber")

		return nil
	}
}
