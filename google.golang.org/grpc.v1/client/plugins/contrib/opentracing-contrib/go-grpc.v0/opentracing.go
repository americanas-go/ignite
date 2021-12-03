package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Register registers a new opentracing plugin for grpc client.
func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewOpentracingWithOptions(o)
	return h.Register(ctx)
}

// Opentracing represents opentracing plugin for grpc client.
type Opentracing struct {
	options *Options
}

// NewOpentracingWithOptions returns a new opentracing plugin with options.
func NewOpentracingWithOptions(options *Options) *Opentracing {
	return &Opentracing{options: options}
}

// NewOpentracingWithConfigPath returns a new opentracing plugin with options from config path.
func NewOpentracingWithConfigPath(path string) (*Opentracing, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOpentracingWithOptions(o), nil
}

// NewOpentracing returns a new opentracing plugin with default options.
func NewOpentracing() *Opentracing {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOpentracingWithOptions(o)
}

// Register registers this opentracing plugin for grpc client.
func (i *Opentracing) Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("opentracing interceptor successfully enabled in grpc client")

	tracer := opentracing.GlobalTracer()

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithChainStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	}, nil

}
