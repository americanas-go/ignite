package promotheus

import hystrix_metric "github.com/afex/hystrix-go/hystrix/metric_collector"

func Register(options *Options) {
	wrapper := NewPrometheusCollector(options.Namespace, options.Labels)
	hystrix_metric.Register(wrapper)
}
