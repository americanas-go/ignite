package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	if !IsEnabled() || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("datadog interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(grpctrace.UnaryClientInterceptor()),
		grpc.WithChainStreamInterceptor(grpctrace.StreamClientInterceptor()),
	}, nil

}
