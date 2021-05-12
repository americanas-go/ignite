package datadog

import (
	"context"

	"github.com/americanas-go/log"
	datadog "github.com/jvitoroc/ignite/datadog/dd-trace-go.v1"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("datadog interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpctrace.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(grpctrace.StreamServerInterceptor()),
	}

}
