package newrelic

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

func Register(ctx context.Context, server *echo.Server) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling newrelic middleware in echo")

	server.Use(nrecho.Middleware(newrelic.Application()))

	logger.Debug("newrelic middleware successfully enabled in echo")

	if IsEnabledRequestID() {

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
