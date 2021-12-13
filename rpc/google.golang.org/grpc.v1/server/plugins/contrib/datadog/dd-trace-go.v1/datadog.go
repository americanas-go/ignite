package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/apm/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

// Register registers a new datadog plugin for grpc server.
func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx)
}

// Datadog represents datadog plugin for grpc server.
type Datadog struct {
	options *Options
}

// NewDatadogWithOptions returns a new datadog plugin with options.
func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

// NewDatadogWithConfigPath returns a new datadog plugin with options from config path.
func NewDatadogWithConfigPath(path string, traceOptions ...grpctrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

// NewDatadog returns a new datadog plugin with default options.
func NewDatadog(traceOptions ...grpctrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

// Register registers this datadog plugin for grpc server.
func (i *Datadog) Register(ctx context.Context) []grpc.ServerOption {

	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("datadog interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpctrace.UnaryServerInterceptor(i.options.traceOptions...)),
		grpc.ChainStreamInterceptor(grpctrace.StreamServerInterceptor(i.options.traceOptions...)),
	}

}
