package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/elastic/go-elasticsearch.v8"
	"github.com/jvitoroc/ignite/elastic/go-elasticsearch.v8/plugins/core/health"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	i := health.NewHealth()

	_, err := elasticsearch.NewClient(context.Background(), i.Register)
	if err != nil {
		log.Panic(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))
}
