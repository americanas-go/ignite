package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	"github.com/americanas-go/health"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/gocql/gocql.v0"
	h "github.com/americanas-go/ignite/gocql/gocql.v0/plugins/contrib/americanas-go/health.v1"
	"github.com/americanas-go/log"
)

func main() {

	config.Load()

	ilog.New()

	i := h.NewHealth()

	session, err := gocql.NewSession(context.Background(), i.Register)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	all := health.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
