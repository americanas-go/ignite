package opentracing

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	apmecho "github.com/opentracing-contrib/echo"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling opentracing middleware in echo")

	server.Use(apmecho.Middleware("echo"))

	logger.Debug("recover opentracing successfully enabled in echo")

	return nil
}
