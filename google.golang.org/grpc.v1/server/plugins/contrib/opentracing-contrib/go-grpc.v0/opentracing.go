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
	h := NewNewrelicWithOptions(o)
	return h.Register(ctx)
}

type Newrelic struct {
	options *Options
}

func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

func NewNewrelic() *Newrelic {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewNewrelicWithOptions(o)
}

func (i *Newrelic) Register(ctx context.Context) []grpc.ServerOption {

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
