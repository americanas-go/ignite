package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/go-redis/redis.v7"
	"github.com/jvitoroc/ignite/go-redis/redis.v7/plugins/core/health"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

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
