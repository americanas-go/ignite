package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/ignite/rpc/google.golang.org/grpc.v1/server"
	"github.com/americanas-go/log"
)

func init() {
	os.Setenv("IGNITE_GRPC_SERVER_TLS_ENABLED", "true")
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

func main() {

	ctx := context.Background()

	config.Load()

	ilog.New()

	srv := server.NewServer(ctx)

	RegisterExampleServer(srv.Server(), &Service{})

	srv.Serve(ctx)
}

type Service struct {
}

func (h *Service) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {

	logger := log.FromContext(ctx)

	logger.Infof(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
