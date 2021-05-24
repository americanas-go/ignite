package swagger

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	eswagger "github.com/swaggo/echo-swagger"
)

func Register(ctx context.Context, server *echo.Server) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	swaggerRoute := GetRoute()

	logger.Tracef("configuring swagger router on %s in echo", swaggerRoute)

	server.GET(swaggerRoute, eswagger.WrapHandler)

	logger.Debugf("swagger router configured on %s in echo", swaggerRoute)

	return nil
}
