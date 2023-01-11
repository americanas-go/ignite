package prometheus

import (
	"context"

	hystrix_metric "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/americanas-go/log"
)

func Register(ctx context.Context) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}

	p := NewPrometheusWithOptions(o)
	return p.Register(ctx)
}

// Prometheus represents a Prometheus plugin for hystrix.
type Prometheus struct {
	options *Options
}

// NewPrometheusWithOptions returns a new Prometheus plugin with options.
func NewPrometheusWithOptions(options *Options) *Prometheus {
	return &Prometheus{options: options}
}

// NewPrometheusWithConfigPath returns a new Prometheus plugin with options from config path.
func NewPrometheusWithConfigPath(path string) (*Prometheus, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPrometheusWithOptions(o), nil
}

// NewPrometheus returns a new Prometheus plugin with default options.
func NewPrometheus() *Prometheus {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewPrometheusWithOptions(o)
}

// Register registers this Prometheus plugin for hystrix.
func (i *Prometheus) Register(ctx context.Context) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling prometheus on hystrix")

	wrapper := NewPrometheusCollector(i.options.Namespace, i.options.Labels)
	hystrix_metric.Registry.Register(wrapper)

	return nil
}
