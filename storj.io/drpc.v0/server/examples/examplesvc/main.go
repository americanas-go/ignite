package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/storj.io/drpc.v0/server"
	"github.com/americanas-go/ignite/storj.io/drpc.v0/server/examples/examplesvc/pb"
	alog "github.com/americanas-go/log"
)

func init() {
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

func main() {

	ctx := context.Background()

	config.Load()

	ilog.New()

	srv, _ := server.NewServer(ctx)
	m := srv.Mux()
	if err := pb.DRPCRegisterExample(m, &Service{}); err != nil {
		panic(err)
	}

	srv.Serve(ctx)
}

type Service struct {
	pb.DRPCExampleUnimplementedServer
}

func (h *Service) Test(ctx context.Context, request *pb.TestRequest) (*pb.TestResponse, error) {

	logger := alog.FromContext(ctx)

	logger.Infof(request.Message)

	return &pb.TestResponse{Message: "hello world"}, nil
}

func NewService() pb.DRPCExampleServer {
	return &Service{}
}
