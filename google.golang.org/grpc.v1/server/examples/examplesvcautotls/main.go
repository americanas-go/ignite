package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server/examples/examplesvc/pb"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server/plugins/contrib/americanas-go/log.v1"
	alog "github.com/americanas-go/log"
)

func init() {
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

func main() {

	ctx := context.Background()

	config.Load()

	ilog.New()

	options, _ := server.NewOptions()
	options.Port = 8080
	options.TLS.Enabled = true
	options.TLS.Auto.Host = "localhost"

	srv := server.NewServerWithOptions(ctx, options, log.Register)

	pb.RegisterExampleServer(srv.ServiceRegistrar(), &Service{})

	srv.Serve(ctx)
}

type Service struct {
	pb.UnimplementedExampleServer
}

func (h *Service) Test(ctx context.Context, request *pb.TestRequest) (*pb.TestResponse, error) {

	logger := alog.FromContext(ctx)

	logger.Infof(request.Message)

	return &pb.TestResponse{Message: "hello world"}, nil
}
