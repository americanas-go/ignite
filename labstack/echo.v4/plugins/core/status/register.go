package status

import (
	"context"
	"net/http"

	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *e.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	statusRoute := GetRoute()

	logger.Tracef("configuring status router on %s in echo", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	logger.Debugf("status router configured on %s in echo", statusRoute)

	return nil
}

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c e.Context) error {
	return echo.JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
}
