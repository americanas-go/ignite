package prometheus

import (
	"context"

	"github.com/americanas-go/log"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

// Register registers a new prometheus plugin for grpc server.
func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewPrometheusWithOptions(o)
	return h.Register(ctx)
}

// Prometheus represents prometheus plugin for grpc server.
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

// Register registers this prometheus plugin for grpc server.
func (i *Prometheus) Register(ctx context.Context) []grpc.ServerOption {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("prometheus interceptor successfully enabled in grpc server")

	buckets := []float64{
		0.000000001, // 1ns
		0.000000002,
		0.000000005,
		0.00000001, // 10ns
		0.00000002,
		0.00000005,
		0.0000001, // 100ns
		0.0000002,
		0.0000005,
		0.000001, // 1µs
		0.000002,
		0.000005,
		0.00001, // 10µs
		0.00002,
		0.00005,
		0.0001, // 100µs
		0.0002,
		0.0005,
		0.001, // 1ms
		0.002,
		0.005,
		0.01, // 10ms
		0.02,
		0.05,
		0.1, // 100 ms
		0.2,
		0.5,
		1.0, // 1s
		2.0,
		5.0,
		10.0, // 10s
		15.0,
		20.0,
		30.0,
	}

	grpc_prometheus.EnableHandlingTimeHistogram(grpc_prometheus.WithHistogramBuckets(buckets))

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.ChainStreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	}

}
