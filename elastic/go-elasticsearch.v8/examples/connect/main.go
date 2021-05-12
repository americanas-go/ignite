package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/elastic/go-elasticsearch.v8"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	client, err := elasticsearch.NewClient(context.Background())
	if err != nil {
		log.Panic(err)
	}

	ping, err := client.Ping(client.Ping.WithPretty())
	if err != nil {
		log.Panic(err)
	}

	log.Infof("status: %v", ping.StatusCode)

}
