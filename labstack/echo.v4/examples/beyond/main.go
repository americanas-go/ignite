package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/core/health"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/core/logger"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/core/status"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/native/cors"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/native/gzip"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/native/requestid"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
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

	logrus.NewLogger()

	srv := echo.NewServer(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		logger.Register,
		status.Register,
		health.Register)

	srv.Instance().GET(c.App.Endpoint.Google, Get)

	srv.Serve(ctx)
}
