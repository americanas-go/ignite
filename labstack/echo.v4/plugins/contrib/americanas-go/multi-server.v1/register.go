package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	checkRoute := getRoute()

	logger.Tracef("configuring multi server check router on %s in echo", checkRoute)

	server.GET(checkRoute, handler)

	logger.Debugf("multi server check router configured on %s in echo", checkRoute)

	return nil
}

func handler(c e.Context) error {

	status := http.StatusOK
	msg := "OK"

	if err := multiserver.Check(c.Request().Context()); err != nil {
		status = http.StatusServiceUnavailable
		msg = "Service Unavailable"
	}

	return c.String(status, msg)
}
