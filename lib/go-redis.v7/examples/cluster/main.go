package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	h "github.com/americanas-go/health"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	iredis "github.com/americanas-go/ignite/lib/go-redis.v7"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
)

func main() {

	config.Load()

	ilog.New()

	var err error

	_, err = iredis.New[*redis.ClusterClient](context.Background())
	if err != nil {
		log.Error(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
