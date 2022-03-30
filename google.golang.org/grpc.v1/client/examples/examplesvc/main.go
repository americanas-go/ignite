package main

import (
	"context"
	"os"
	"time"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/client"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/client/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server/examples/examplesvc/pb"
	alog "github.com/americanas-go/log"
)

func init() {
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

func main() {

	ctx := context.Background()

	config.Load()

	ilog.New()

	request := &pb.TestRequest{
		Message: "mensagem da requisição",
	}

	options, _ := client.NewOptions()
	options.Host = "localhost"
	options.Port = 8080
	options.TLS.Enabled = true
	options.TLS.InsecureSkipVerify = true

	conn := client.NewClientConnWithOptions(ctx, options, log.Register)
	defer conn.Close()

	c := pb.NewExampleClient(conn)

	rctx, _ := context.WithTimeout(ctx, 1*time.Minute)

	test, err := c.Test(rctx, request)
	if err != nil {
		alog.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	alog.Infof(test.Message)

	alog.Infof(conn.GetState().String())
}
