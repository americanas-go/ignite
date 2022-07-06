package json

import (
	"context"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	"github.com/goccy/go-json"
	f "github.com/gofiber/fiber/v2"
)

// Register registers a new goccy/go-json plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewJsonWithOptions(o)
	return h.Register(ctx, options)
}

// Json represents a goccy/go-json plugin for fiber.
type Json struct {
	options *Options
}

// NewJsonWithOptions returns a new goccy/go-json plugin with options.
func NewJsonWithOptions(options *Options) *Json {
	return &Json{options: options}
}

// NewJsonWithConfigPath returns a new goccy/go-json plugin with options from config path.
func NewJsonWithConfigPath(path string) (*Json, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewJsonWithOptions(o), nil
}

// NewJson returns a new goccy/go-json plugin with default options.
func NewJson() *Json {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewJsonWithOptions(o)
}

// Register registers this goccy/go-json plugin for fiber.
func (i *Json) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling goccy/go-json encoder in fiber")

	return func(ctx context.Context, config *f.Config) error {
		config.JSONEncoder = json.Marshal
		config.JSONDecoder = json.Unmarshal
		return nil
	}, nil
}
