package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	logger "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/extra/error_handler"
	"github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/native/cors"
	"github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/native/etag"
	f "github.com/gofiber/fiber/v2"
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

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(c *f.Ctx) (err error) {

	resp := Response{
		Message: "Hello World!!",
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logger.New()

	handler := &Handler{}

	srv := fiber.NewServer(ctx,
		error_handler.Register,
		cors.Register,
		etag.Register)

	srv.Get(c.App.Endpoint.Helloworld, handler.Get)
	srv.Serve(ctx)
}
