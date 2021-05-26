package main

import (
	"context"
	"net/http"
	"os"

	"github.com/americanas-go/config"
	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	logplugin "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/americanas-go/log.v1"
	prometheus "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/prometheus/client_golang.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/extra/error_handler"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/native/cors"
	e "github.com/labstack/echo/v4"
)

func errorHandler(c e.Context) (err error) {
	return errors.NotFoundf("example")
}

func helloHandler(c e.Context) (err error) {
	c.String(http.StatusOK, "hello world")
	return nil
}

func main() {

	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")

	config.Load()
	log.New()

	ctx := context.Background()

	srv := echo.NewServer(ctx,
		cors.Register,
		logplugin.Register,
		prometheus.Register,
		error_handler.Register)

	srv.GET("/not-found", errorHandler)
	srv.GET("/hello", helloHandler)

	srv.Serve(ctx)
}
