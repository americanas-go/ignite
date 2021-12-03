package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Register registers a new opentracing plugin for grpc server.
func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewOpenTracingWithOptions(o)
	return h.Register(ctx)
}

// OpenTracing represents opentracing plugin for grpc server.
type OpenTracing struct {
	options *Options
}

// NewOpenTracingWithOptions returns a new opentracing plugin with options.
func NewOpenTracingWithOptions(options *Options) *OpenTracing {
	return &OpenTracing{options: options}
}

// NewOpenTracingWithConfigPath returns a new opentracing plugin with options from config path.
func NewOpenTracingWithConfigPath(path string) (*OpenTracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpenTracingWithOptions(o), nil
}

// NewOpenTracing returns a new opentracing plugin with default options.
func NewOpenTracing() *OpenTracing {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOpenTracingWithOptions(o)
}

// Register registers this opentracing plugin for grpc server.
func (i *OpenTracing) Register(ctx context.Context) []grpc.ServerOption {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("opentracing interceptor successfully enabled in grpc server")

	tracer := opentracing.GlobalTracer()

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.ChainStreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	}

}
