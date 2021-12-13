package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	ilog.New()

	conn, err := mongo.NewConn(context.Background(), newrelic.Register)
	if err != nil {
		log.Panic(err)
	}

	err = conn.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

}
