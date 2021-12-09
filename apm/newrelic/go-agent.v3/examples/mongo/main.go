package main

import (
	"context"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
)

func main() {

	config.Load()
	ilog.New()

	mongo.NewConn(context.Background(), newrelic.Register)
}
