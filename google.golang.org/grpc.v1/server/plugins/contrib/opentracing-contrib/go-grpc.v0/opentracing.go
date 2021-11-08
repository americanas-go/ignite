package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewOpenTracingWithOptions(o)
	return h.Register(ctx)
}

type OpenTracing struct {
	options *Options
}

func NewOpenTracingWithOptions(options *Options) *OpenTracing {
	return &OpenTracing{options: options}
}

func NewOpenTracingWithConfigPath(path string) (*OpenTracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpenTracingWithOptions(o), nil
}

func NewOpenTracing() *OpenTracing {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOpenTracingWithOptions(o)
}

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
