package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/google.golang.org/grpc.v1/server"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func main() {

	ctx := context.Background()

	config.Load()

	// start logrus
	logrus.NewLogger()

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
