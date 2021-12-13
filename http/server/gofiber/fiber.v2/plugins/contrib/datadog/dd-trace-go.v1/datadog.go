package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/apm/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	f "github.com/gofiber/fiber/v2"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

// Register registers a new datadog plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx, options)
}

// Datadog represents a datadog plugin for fiber.
type Datadog struct {
	options *Options
}

// NewDatadogWithOptions returns a new datadog plugin with options.
func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

// NewDatadogWithConfigPath returns a new datadog plugin with options from config path.
func NewDatadogWithConfigPath(path string, traceOptions ...fibertrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

// NewDatadog returns a new datadog plugin with default options.
func NewDatadog(traceOptions ...fibertrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

// Register registers this datadog plugin for fiber.
func (i *Datadog) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in fiber")

	return nil, func(ctx context.Context, app *f.App) error {

		app.Use(fibertrace.Middleware(i.options.TraceOptions...))

		logger.Debug("datadog middleware successfully enabled in fiber")

		return nil

	}
}
