package main

import (
	"context"
	"net/http"
	"os"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
)

func init() {
	fiber.ConfigAdd("ignite.fiber2")
}

type Response struct {
	Message string
}

func Get(c *f.Ctx) (err error) {

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}

func main() {

	os.Setenv("IGNITE_FIBER2_PORT", "8086")

	config.Load()

	ilog.New()

	ctx1 := context.Background()

	srv1 := fiber.NewServer(ctx1)

	srv1.Get("/test", Get)

	ctx2 := context.Background()

	srv2, err := fiber.NewServerWithConfigPath(ctx2, "ignite.fiber2")
	if err != nil {
		panic(err)
	}

	srv2.Get("/test", Get)

	multiserver.Serve(context.Background(), srv1, srv2)
}
