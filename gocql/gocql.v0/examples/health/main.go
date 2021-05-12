package main

import (
	"context"
	"encoding/json"

	"github.com/americanas-go/config"
	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/gocql/gocql.v0"
	h "github.com/jvitoroc/ignite/gocql/gocql.v0/plugins/core/health"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

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
