package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/client/go-resty/resty.v2"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/contrib/americanas-go/health.v1"
	logplugin "github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/contrib/americanas-go/log.v1"
	status "github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/contrib/americanas-go/rest-response.v1"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/native/cors"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/native/gzip"
	"github.com/americanas-go/ignite/http/server/labstack/echo.v4/plugins/native/requestid"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	r "github.com/go-resty/resty/v2"
	e "github.com/labstack/echo/v4"
)

const Endpoint = "app.endpoint.google"

func init() {
	config.Add(Endpoint, "/google", "google endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Google string
		}
	}
}

type Response struct {
	Message string
}

type Handler struct {
	client *r.Client
}

func NewHandler(client *r.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Get(c e.Context) (err error) {

	request := h.client.R().EnableTrace()

	_, err = request.Get("http://google.com")
	if err != nil {
		return err
	}

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	ilog.New()

	srv := echo.NewServer(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		logplugin.Register,
		status.Register,
		health.Register)

	// instance.AddErrorAdvice(customErrors.InvalidPayload, 400)

	options := resty.Options{
		Host: "http://www.google.com",
	}

	client := resty.NewClientWithOptions(ctx, &options)

	handler := NewHandler(client)
	srv.GET(c.App.Endpoint.Google, handler.Get)

	srv.Serve(ctx)
}
