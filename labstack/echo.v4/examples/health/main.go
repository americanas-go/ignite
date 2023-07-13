package main

import (
	"context"
	"github.com/americanas-go/health"
	"net/http"
	"os"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	h "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/americanas-go/health.v1"
	e "github.com/labstack/echo/v4"
)

func helloHandler(c e.Context) (err error) {
	c.String(http.StatusOK, "hello world")
	return nil
}

type MyChecker struct {
}

func (c *MyChecker) Check(ctx context.Context) error {
	return nil
}

func main() {

	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")

	config.Load()
	log.New()

	ctx := context.Background()

	hc := health.NewHealthChecker("teste", "teste", &MyChecker{}, true, true)
	health.Add(hc)

	srv := echo.NewServer(ctx, h.Register)

	srv.GET("/hello", helloHandler)

	srv.Serve(ctx)
}
