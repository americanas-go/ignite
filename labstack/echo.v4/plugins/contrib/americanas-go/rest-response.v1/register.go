package status

import (
	"context"
	"net/http"

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

	statusRoute := GetRoute()

	logger.Tracef("configuring status router on %s in echo", statusRoute)

	server.GET(statusRoute, handler)

	logger.Debugf("status router configured on %s in echo", statusRoute)

	return nil
}

func handler(c e.Context) error {
	return c.JSON(http.StatusOK, response.NewResourceStatus())
}
