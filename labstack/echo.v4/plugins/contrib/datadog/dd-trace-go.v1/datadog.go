package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx, server)
}

type Datadog struct {
	options *Options
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func NewDatadogWithConfigPath(path string, traceOptions ...echotrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadog(traceOptions ...echotrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return NewDatadogWithOptions(o)
}

func (i *Datadog) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling datadog middleware in echo")

	server.Use(echotrace.Middleware(i.options.TraceOptions...))

	logger.Debug("datadog middleware successfully enabled in echo")

	return nil
}
