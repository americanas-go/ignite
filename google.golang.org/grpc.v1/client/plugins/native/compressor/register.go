package compressor

import (
	"context"

	"github.com/americanas-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	logger := log.FromContext(ctx)
	logger.Debug("compressor successfully enabled in grpc client")

	err := gzip.SetLevel(Level())
	if err != nil {
		logger.Fatalf("could not set level: %s", err.Error())
	}

	return nil, []grpc.CallOption{
		grpc.UseCompressor(gzip.Name),
	}
}
