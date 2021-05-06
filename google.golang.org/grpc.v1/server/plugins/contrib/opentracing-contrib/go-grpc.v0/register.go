package opentracing

import (
	"context"

	"github.com/americanas-go/log"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() {
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
