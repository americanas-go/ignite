package cors

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

	logger.Trace("enabling cors middleware in echo")

	instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     GetAllowOrigins(),
		AllowMethods:     GetAllowMethods(),
		AllowHeaders:     GetAllowHeaders(),
		AllowCredentials: GetAllowCredentials(),
		ExposeHeaders:    GetExposeHeaders(),
		MaxAge:           GetMaxAge(),
	}))

	logger.Debug("cors middleware successfully enabled in echo")

	return nil
}
