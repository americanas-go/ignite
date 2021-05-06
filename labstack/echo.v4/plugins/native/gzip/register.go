package gzip

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling gzip middleware in echo")

	instance.Use(middleware.Gzip())

	logger.Debug("gzip middleware successfully enabled in echo")

	return nil
}
