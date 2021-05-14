package main

import (
	"context"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/client"
	"github.com/americanas-go/log"
)

func main() {

	ctx := context.Background()

	config.Load()

	ilog.New()

	request := &TestRequest{
		Message: "mensagem da requisição",
	}

	options := client.Options{
		Host: "localhost",
		Port: 9090,
	}

	conn := client.NewClientConnWithOptions(ctx, &options)
	defer conn.Close()

	c := NewExampleClient(conn)

	test, err := c.Test(ctx, request)
	if err != nil {
		log.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	log.Infof(test.Message)

	log.Infof(conn.GetState().String())
}
