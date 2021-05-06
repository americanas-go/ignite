package fiber

import (
	"context"
	"strconv"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
)

type Plugin func(context.Context, *fiber.App) error

type Server struct {
	app     *fiber.App
	options *Options
}

func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, options, plugins...)
}

func NewServerWithOptions(ctx context.Context, options *Options, plugins ...Plugin) *Server {

	app := fiber.New(*options.Config)

	for _, plugin := range plugins {
		if err := plugin(ctx, app); err != nil {
			panic(err)
		}
	}

	return &Server{app: app, options: options}
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)
	logger.Infof("starting fiber Server. https://gofiber.io/")

	addr := ":" + strconv.Itoa(s.options.Port)

	logger.Fatal(s.app.Listen(addr))
}
