package main

import (
	"context"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/americanas-go/health.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/americanas-go/log.v1"
	status "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/americanas-go/rest-response.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/native/cors"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/native/gzip"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/native/requestid"
	"github.com/wesovilabs/beyond/api"
)

const Endpoint = "app.endpoint.google"

func init() {
	config.Add(Endpoint, "/google", "google endpoint")
}

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(NewTracingAdvice, "handler.Get(...)").
		WithBefore(NewTracingAdviceWithPrefix("[beyond]"), "handler.*(...)...")
}

func main() {

	var err error

	config.Load()

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	ilog.New()

	srv := echo.NewServer(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		log.Register,
		status.Register,
		health.Register)

	srv.GET(c.App.Endpoint.Google, Get)

	srv.Serve(ctx)
}
