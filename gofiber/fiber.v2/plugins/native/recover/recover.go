package recover

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewRecoverWithOptions(o)
	return n.Register(ctx, options)
}

type Recover struct {
	options *Options
}

func NewRecoverWithConfigPath(path string) (*Recover, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRecoverWithOptions(o), nil
}

func NewRecoverWithOptions(options *Options) *Recover {
	return &Recover{options: options}
}

func (d *Recover) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling recover middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(recover.New())

		logger.Debug("recover middleware successfully enabled in fiber")

		return nil
	}
}
