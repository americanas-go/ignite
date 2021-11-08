package compress

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewCompressWithOptions(o)
	return n.Register(ctx, options)
}

type Compress struct {
	options *Options
}

func NewCompressWithConfigPath(path string) (*Compress, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCompressWithOptions(o), nil
}

func NewCompressWithOptions(options *Options) *Compress {
	return &Compress{options: options}
}

func (d *Compress) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling compress middleware in fiber")

	lvl := d.options.GetLevel()

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(compress.New(compress.Config{
			Level: lvl,
		}))

		logger.Debug("compress middleware successfully enabled in fiber")

		return nil
	}
}
