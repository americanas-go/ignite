package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
)

func main() {

	config.Load()
	ilog.New()

	mongo.NewConn(context.Background(), newrelic.Register)
}
