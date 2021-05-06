package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	"github.com/americanas-go/ignite/go-redis/redis.v7"
	"github.com/americanas-go/ignite/go-redis/redis.v7/plugins/core/health"
	"github.com/americanas-go/ignite/sirupsen/logrus.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	logrus.NewLogger()

	var err error

	healthIntegrator := health.NewClientHealth()

	_, err = redis.NewClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		log.Error(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
