package recover

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling recover middleware in echo")

	instance.Use(middleware.Recover())

	logger.Debug("recover middleware successfully enabled in echo")

	return nil
}
