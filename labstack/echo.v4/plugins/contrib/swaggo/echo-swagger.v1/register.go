package swagger

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
	eswagger "github.com/swaggo/echo-swagger"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	swaggerRoute := GetRoute()

	logger.Tracef("configuring swagger router on %s in echo", swaggerRoute)

	instance.GET(swaggerRoute, eswagger.WrapHandler)

	logger.Debugf("swagger router configured on %s in echo", swaggerRoute)

	return nil
}
