package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
	apmecho "github.com/opentracing-contrib/echo"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling opentracing middleware in echo")

	instance.Use(apmecho.Middleware("echo"))

	logger.Debug("recover opentracing successfully enabled in echo")

	return nil
}
