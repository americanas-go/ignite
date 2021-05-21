package fiber

import (
	"context"
	"strconv"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
)

type ConfigPlugin func(context.Context, *fiber.Config) error
type AppPlugin func(context.Context, *fiber.App) error

type Plugin func(context.Context, *Options) (ConfigPlugin, AppPlugin)

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

	logger := log.FromContext(ctx)

	var configPlugins []ConfigPlugin
	var appPlugins []AppPlugin

	for _, plugin := range plugins {
		configPlugin, appPlugin := plugin(ctx, options)
		if configPlugin != nil {
			configPlugins = append(configPlugins, configPlugin)
		}
		if appPlugin != nil {
			appPlugins = append(appPlugins, appPlugin)
		}
	}

	for _, configPlugin := range configPlugins {
		if err := configPlugin(ctx, options.Config); err != nil {
			logger.Fatalf(err.Error())
		}
	}

	app := fiber.New(*options.Config)

	for _, appPlugin := range appPlugins {
		if err := appPlugin(ctx, app); err != nil {
			panic(err)
		}
	}

	return &Server{app: app, options: options}
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) Get(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Get(path, handlers...)
}

func (s *Server) Post(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Post(path, handlers...)
}

func (s *Server) Delete(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Delete(path, handlers...)
}

func (s *Server) Head(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Head(path, handlers...)
}

func (s *Server) Put(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Put(path, handlers...)
}

func (s *Server) Patch(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Patch(path, handlers...)
}

func (s *Server) Add(method string, path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Add(method, path, handlers...)
}

func (s *Server) Connect(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Connect(path, handlers...)
}

func (s *Server) All(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.All(path, handlers...)
}

func (s *Server) Group(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Group(path, handlers...)
}

func (s *Server) Options(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Options(path, handlers...)
}

func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)
	logger.Infof("starting fiber Server. https://gofiber.io/")

	addr := ":" + strconv.Itoa(s.options.Port)

	logger.Error(s.app.Listen(addr))
}

func (s *Server) Shutdown(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Error(s.app.Shutdown())
}
