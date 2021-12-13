package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

// Register registers a new newrelic plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewNewrelicWithOptions(o)
	return h.Register(ctx, server)
}

// Newrelic represents newrelic plugin for echo server.
type Newrelic struct {
	options *Options
}

// NewNewrelicWithOptions returns a new newrelic plugin with options.
func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

// NewNewrelicWithConfigPath returns a new newrelic plugin with options from config path.
func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

// NewNewrelic returns a new newrelic plugin with default options.
func NewNewrelic() *Newrelic {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewNewrelicWithOptions(o)
}

// Register registers this newrelic plugin for echo server.
func (i *Newrelic) Register(ctx context.Context, server *echo.Server) error {

	if !i.options.Enabled || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling newrelic middleware in echo")

	server.Use(nrecho.Middleware(newrelic.Application()))

	logger.Debug("newrelic middleware successfully enabled in echo")

	if i.options.Middlewares.RequestID.Enabled {

		logger.Trace("enabling requestID newrelic middleware in echo")

		server.Use(requestIDMiddleware())

		logger.Debug("requestID newrelic middleware successfully enabled in echo")
	}

	return nil
}

func requestIDMiddleware() e.MiddlewareFunc {
	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) (err error) {
			ctx := c.Request().Context()
			txn := newrelic.FromContext(ctx)
			reqId := c.Request().Header.Get(e.HeaderXRequestID)
			if reqId == "" {
				reqId = c.Response().Header().Get(e.HeaderXRequestID)
			}

			txn.AddAttribute("request.id", reqId)
			return next(c)
		}
	}
}
