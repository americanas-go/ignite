package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	if !IsEnabled() {
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
