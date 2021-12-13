package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	"github.com/americanas-go/ignite/db/nosql/go-redis/redis.v7"
	"github.com/americanas-go/ignite/db/nosql/go-redis/redis.v7/plugins/contrib/americanas-go/health.v1"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	ilog.New()

	var err error

	healthIntegrator := health.NewClusterHealth()

	_, err = redis.NewClusterClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		log.Error(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
