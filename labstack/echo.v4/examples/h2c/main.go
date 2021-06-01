package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	logplugin "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/americanas-go/log.v1"
	"github.com/americanas-go/ignite/labstack/echo.v4/plugins/extra/error_handler"
	e "github.com/labstack/echo/v4"
)

func init() {
	os.Setenv("IGNITE_ECHO_PROTOCOL", "H2C")
	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "TRACE")
}

func main() {

	config.Load()
	log.New()

	ctx := context.Background()

	srv := echo.NewServer(ctx,
		logplugin.Register,
		error_handler.Register)

	srv.GET("/", func(c e.Context) error {
		req := c.Request()
		format := `
			<code>
			  Protocol: %s<br>
			  Host: %s<br>
			  Remote Address: %s<br>
			  Method: %s<br>
			  Path: %s<br>
			</code>
		  `
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})

	srv.Serve(ctx)

	// curl -v --http2-prior-knowledge http://localhost:8080
}
