package health

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	healthRoute := GetRoute()

	logger.Tracef("configuring health router on %s in echo", healthRoute)

	server.GET(healthRoute, handler)

	logger.Debugf("health router configured on %s in echo", healthRoute)

	return nil
}

func handler(c e.Context) error {

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	resp, httpCode := response.NewHealth(ctx)

	return c.JSON(httpCode, resp)
}
