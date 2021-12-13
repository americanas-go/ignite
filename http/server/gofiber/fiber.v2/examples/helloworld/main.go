package main

import (
	"context"
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2/plugins/extra/error_handler"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2/plugins/native/cors"
	"github.com/americanas-go/ignite/http/server/gofiber/fiber.v2/plugins/native/etag"
	"github.com/americanas-go/ignite/log/americanas-go/log.v1"
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
