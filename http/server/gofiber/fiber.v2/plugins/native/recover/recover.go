package recover

import (
	"context"

	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Register registers a new recover plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewRecoverWithOptions(o)
	return n.Register(ctx, options)
}

// Recover represents a new recover plugin for fiber.
type Recover struct {
	options *Options
}

// NewRecoverWithConfigPath returns a new recover plugin with options from config path.
func NewRecoverWithConfigPath(path string) (*Recover, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRecoverWithOptions(o), nil
}

// NewRecoverWithOptions returns a new recover plugin with options.
func NewRecoverWithOptions(options *Options) *Recover {
	return &Recover{options: options}
}

// Register registers this recover plugin for fiber.
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
