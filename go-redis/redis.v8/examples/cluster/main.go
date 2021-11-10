package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/go-redis/redis.v8"
	"github.com/americanas-go/ignite/go-redis/redis.v8/plugins/contrib/americanas-go/health.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	ilog.New()

	var err error

	healthIntegrator, err := health.NewClusterHealth()
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = redis.NewClusterClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		log.Error(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
