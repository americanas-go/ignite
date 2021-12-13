package opentracing

import (
	"context"

	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
	"github.com/americanas-go/log"
	apmecho "github.com/opentracing-contrib/echo"
)

// Register registers a new opentracing plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewOpenTracingWithOptions(o)
	return h.Register(ctx, server)
}

// OpenTracing represents opentracing plugin for echo server.
type OpenTracing struct {
	options *Options
}

// NewOpenTracingWithOptions returns a new opentracing plugin with options.
func NewOpenTracingWithOptions(options *Options) *OpenTracing {
	return &OpenTracing{options: options}
}

// NewOpenTracingWithConfigPath returns a new opentracing plugin with options from config path.
func NewOpenTracingWithConfigPath(path string) (*OpenTracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpenTracingWithOptions(o), nil
}

// NewOpenTracing returns a new opentracing plugin with default options.
func NewOpenTracing() *OpenTracing {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOpenTracingWithOptions(o)
}

// Register registers this opentracing plugin for echo server.
func (i *OpenTracing) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling opentracing middleware in echo")

	server.Use(apmecho.Middleware("echo"))

	logger.Debug("recover opentracing successfully enabled in echo")

	return nil
}
