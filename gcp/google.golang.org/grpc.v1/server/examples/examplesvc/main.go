package main

import (
	"context"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
	"github.com/americanas-go/log"
)

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
