package requestid

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling requestID middleware in echo")

	server.Use(middleware.RequestID())

	logger.Debug("requestID middleware successfully enabled in echo")

	return nil
}
