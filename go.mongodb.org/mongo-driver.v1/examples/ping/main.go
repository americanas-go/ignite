package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	newrelic "github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1/plugins/contrib/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	logrus.NewLogger()

	conn, err := mongo.NewConn(context.Background(), newrelic.Register)
	if err != nil {
		log.Panic(err)
	}

	err = conn.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

}
