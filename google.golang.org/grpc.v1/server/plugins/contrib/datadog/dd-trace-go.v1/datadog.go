package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx)
}

type Datadog struct {
	options *Options
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func NewDatadogWithConfigPath(path string, traceOptions ...grpctrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadog(traceOptions ...grpctrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

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