package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/core/health"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/core/logger"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/core/status"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/native/cors"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/native/gzip"
	"github.com/jvitoroc/ignite/labstack/echo.v4/plugins/native/requestid"
	"github.com/jvitoroc/ignite/sirupsen/logrus.v1"
	e "github.com/labstack/echo/v4"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Helloworld string
		}
	}
}

type Response struct {
	Message string
}

func Get(c e.Context) (err error) {

	l := log.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return echo.JSON(c, http.StatusOK, resp, err)
}

func main() {

	config.Load()
	logrus.NewLogger()
	//zap.NewLogger()
	//zerolog.NewLogger()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	srv := echo.NewServer(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		logger.Register,
		status.Register,
		health.Register)

	srv.Instance().GET(c.App.Endpoint.Helloworld, Get)

	srv.Serve(ctx)
}
