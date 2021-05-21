package log

import (
	"context"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("logger interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainStreamInterceptor(streamInterceptor()),
		grpc.WithChainUnaryInterceptor(unaryInterceptor()),
	}, nil
}
