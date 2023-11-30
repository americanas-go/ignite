package contrib // import "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/opentelemetry/otelecho.v1"

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

// Register registers a new opentelemetry plugin for echo server.
func Register(ctx context.Context, server *echo.Server, tracingOptions ...otelecho.Option) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	o.TracingOptions = tracingOptions
	h := NewOtelEchoWithOptions(o)
	h.Register(ctx, server)
	return nil
}

// OtelEcho represents opentelemetry plugin for echo server.
type OtelEcho struct {
	options *Options
}

// NewOtelEchoWithOptions returns a new opentelemetry plugin with options.
func NewOtelEchoWithOptions(options *Options) *OtelEcho {
	return &OtelEcho{options: options}
}

// NewOtelEchoWithConfigPath returns a new opentelemetry plugin with options from config path.
func NewOtelEchoWithConfigPath(path string, tracingOptions ...otelecho.Option) (*OtelEcho, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	o.TracingOptions = tracingOptions
	return NewOtelEchoWithOptions(o), nil
}

// NewOtelEcho returns a new opentelemetry plugin with default options.
func NewOtelEcho(tracingOptions ...otelecho.Option) *OtelEcho {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	o.TracingOptions = tracingOptions
	return NewOtelEchoWithOptions(o)
}

// Register registers this opentelemetry plugin for echo server.
func (i *OtelEcho) Register(ctx context.Context, server *echo.Server) {
	if !i.options.Enabled {
		return
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling opentelemetry middleware in echo")

	server.Use(otelecho.Middleware("", i.options.TracingOptions...))

	logger.Debug("opentelemetry integration successfully enabled in echo")
}
