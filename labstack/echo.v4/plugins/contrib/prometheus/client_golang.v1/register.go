package prometheus

import (
	"context"

	"github.com/americanas-go/log"
	prometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling prometheus middleware in echo")

	instance.Use(prometheus.MetricsMiddleware())

	logger.Debug("prometheus middleware successfully enabled in echo")

	prometheusRoute := GetRoute()

	logger.Tracef("configuring prometheus metric router on %s in echo", prometheusRoute)

	instance.GET(prometheusRoute, echo.WrapHandler(promhttp.Handler()))

	logger.Debugf("prometheus metric router configured on %s in echo", prometheusRoute)

	return nil
}
