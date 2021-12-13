package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/gocql/gocql.v0"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
)

func main() {

	config.Load()

	ilog.New()

	session, err := gocql.NewSession(context.Background())
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.Query("void").Exec()
	if err != nil {
		panic(err)
	}

}
