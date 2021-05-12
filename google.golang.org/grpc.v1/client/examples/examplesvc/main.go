package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/google.golang.org/grpc.v1/client"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func main() {

	ctx := context.Background()

	config.Load()

	logger := logrus.NewLogger()

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
		logger.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	logger.Infof(test.Message)

	logger.Infof(conn.GetState().String())
}
