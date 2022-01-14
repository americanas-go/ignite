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

	c, err := iredis.New[*redis.Client](context.Background())
	if err != nil {
		log.Error(err)
	}
	ping := c.Conn().Ping()
	if ping.Err() != nil {
		log.Error(ping.Err())
	}
	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
