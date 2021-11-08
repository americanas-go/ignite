package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/elastic/go-elasticsearch.v8"
	"github.com/americanas-go/ignite/elastic/go-elasticsearch.v8/plugins/contrib/americanas-go/health.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	ilog.New()

	_, err := elasticsearch.NewClient(context.Background(), health.Register)
	if err != nil {
		log.Panic(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))
}
