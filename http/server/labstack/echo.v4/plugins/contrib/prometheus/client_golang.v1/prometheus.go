package prometheus

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	prometheus "github.com/globocom/echo-prometheus"
	e "github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Register registers a new prometheus plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewPrometheusWithOptions(o)
	return h.Register(ctx, server)
}

// Prometheus represents prometheus plugin for echo server.
type Prometheus struct {
	options *Options
}

// NewPrometheusWithOptions returns a new prometheus plugin with options.
func NewPrometheusWithOptions(options *Options) *Prometheus {
	return &Prometheus{options: options}
}

// NewPrometheusWithConfigPath returns a new prometheus plugin with options from config path.
func NewPrometheusWithConfigPath(path string) (*Prometheus, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPrometheusWithOptions(o), nil
}

// NewPrometheus returns a new prometheus plugin with default options.
func NewPrometheus() *Prometheus {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewPrometheusWithOptions(o)
}

// Register registers this prometheus plugin for echo server.
func (i *Prometheus) Register(ctx context.Context, server *echo.Server) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling prometheus middleware in echo")

	server.Use(prometheus.MetricsMiddleware())

	logger.Debug("prometheus middleware successfully enabled in echo")

	prometheusRoute := i.options.Route

	logger.Tracef("configuring prometheus metric router on %s in echo", prometheusRoute)

	server.GET(prometheusRoute, e.WrapHandler(promhttp.Handler()))

	logger.Debugf("prometheus metric router configured on %s in echo", prometheusRoute)

	return nil
}
