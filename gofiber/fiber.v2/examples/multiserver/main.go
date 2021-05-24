package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	mserver "github.com/americanas-go/ignite/gofiber/fiber.v2/plugins/contrib/americanas-go/multi-server.v1"
	"github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
)

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

	config.Load()

	ilog.New()

	ctx1 := context.Background()

	srv1 := fiber.NewServer(ctx1,
		mserver.Register)

	srv1.Get("/test", Get)

	multiserver.Serve(context.Background(), srv1, &LocalServer{})
}

type LocalServer struct {
}

func (s *LocalServer) Serve(ctx context.Context) {
	time.Sleep(10 * time.Second)
	fmt.Printf("finished")
}

func (s *LocalServer) Shutdown(ctx context.Context) {
}
