package prometheus

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prometheus struct {
	options *Options
}

func NewPrometheusWithConfigPath(path string) (*Prometheus, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPrometheusWithOptions(o), nil
}

func NewPrometheusWithOptions(options *Options) *Prometheus {
	return &Prometheus{options: options}
}

func (d *Prometheus) Register(ctx context.Context) (*chi.Config, error) {

	if d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling prometheus middleware in chi")

	prometheusRoute := d.options.Route

	logger.Tracef("configuring prometheus router on %s in chi", prometheusRoute)

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			promMiddleware,
		},
		Handlers: []chi.ConfigHandler{
			{
				Handler: promhttp.Handler(),
				Pattern: prometheusRoute,
			},
		},
	}, nil

}

func Register(ctx context.Context) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewPrometheusWithOptions(o)
	return n.Register(ctx)
}
