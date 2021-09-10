package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx, options)
}

type Datadog struct {
	options *Options
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func NewDatadogWithConfigPath(path string) (*Datadog, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadog() *Datadog {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

func (i *Datadog) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {

		app.Use(fibertrace.Middleware(i.options.Options...))

		logger.Debug("datadog middleware successfully enabled in fiber")

		return nil

	}
}
