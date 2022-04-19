package main

import (
	"context"
	"os"
	"time"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/storj.io/drpc.v0/client"
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

	request := &pb.TestRequest{
		Message: "mensagem da requisição",
	}

	conn, _ := client.NewClientConn(ctx)
	defer conn.Close()

	c := pb.NewDRPCExampleClient(conn)

	rctx, _ := context.WithTimeout(ctx, 1*time.Minute)

	test, err := c.Test(rctx, request)
	if err != nil {
		alog.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	alog.Infof(test.Message)
}
