package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/contrib/americanas-go/health.v1"
	logplugin "github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/contrib/americanas-go/log.v1"
	status "github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/contrib/americanas-go/rest-response.v1"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/native/cors"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/native/gzip"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/native/requestid"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
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

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {

	config.Load()

	ilog.New()

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
		logplugin.Register,
		status.Register,
		health.Register)

	srv.GET(c.App.Endpoint.Helloworld, Get)

	srv.Serve(ctx)
}
