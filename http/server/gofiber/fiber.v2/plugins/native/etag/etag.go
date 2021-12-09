package etag

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

// Register registers a new etag plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewETagWithOptions(o)
	return n.Register(ctx, options)
}

// ETag represents a new etag plugin for fiber.
type ETag struct {
	options *Options
}

// NewETagWithConfigPath returns a new etag plugin with options from config path.
func NewETagWithConfigPath(path string) (*ETag, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewETagWithOptions(o), nil
}

// NewETagWithOptions returns a new etag plugin with options.
func NewETagWithOptions(options *Options) *ETag {
	return &ETag{options: options}
}

// Register registers this etag plugin for fiber.
func (d *ETag) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling etag middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {
		app.Use(etag.New())

		logger.Debug("etag middleware successfully enabled in fiber")

		return nil
	}
}
