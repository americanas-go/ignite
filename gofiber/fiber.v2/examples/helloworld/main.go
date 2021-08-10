package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/extra/error_handler"
	"github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/native/cors"
	"github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/native/etag"
	f "github.com/gofiber/fiber/v2"
)

func Get(c *f.Ctx) (err error) {
	return c.Status(http.StatusOK).SendString("Hello!!")
}

func main() {

	config.Load()
	log.New()

	ctx := context.Background()

	srv := fiber.NewServer(ctx,
		error_handler.Register,
		cors.Register,
		etag.Register)

	srv.Get("/hello-world", Get)
	srv.Serve(ctx)
}
