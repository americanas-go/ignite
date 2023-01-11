package json

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	"github.com/bytedance/sonic"
	f "github.com/gofiber/fiber/v2"
)

// Register registers a new sonic plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewSonicWithOptions(o)
	return h.Register(ctx, options)
}

// Sonic represents a sonic plugin for fiber.
type Sonic struct {
	options *Options
}

// NewSonicWithOptions returns a new sonic plugin with options.
func NewSonicWithOptions(options *Options) *Sonic {
	return &Sonic{options: options}
}

// NewSonicWithConfigPath returns a new sonic plugin with options from config path.
func NewSonicWithConfigPath(path string) (*Sonic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSonicWithOptions(o), nil
}

// NewSonic returns a new sonic plugin with default options.
func NewSonic() *Sonic {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewSonicWithOptions(o)
}

// Register registers this sonic plugin for fiber.
func (i *Sonic) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling bytedance/sonic encoder in fiber")

	return func(ctx context.Context, config *f.Config) error {
		config.JSONEncoder = sonic.Marshal
		config.JSONDecoder = sonic.Unmarshal
		return nil
	}, nil
}
