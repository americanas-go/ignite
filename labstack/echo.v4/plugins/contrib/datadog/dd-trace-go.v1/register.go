package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	ddecho "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling datadog middleware in echo")

	server.Use(ddecho.Middleware(
		ddecho.WithServiceName(datadog.Service()),
		ddecho.WithAnalyticsRate(datadog.AnalyticsRate()),
	))

	logger.Debug("datadog middleware successfully enabled in echo")

	return nil
}
