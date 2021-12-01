package fiber

import (
	"context"
	"strconv"

	"github.com/americanas-go/log"
	"github.com/gofiber/fiber/v2"
)

// ConfigPlugin defines a ConfigPlugin function signature.
type ConfigPlugin func(context.Context, *fiber.Config) error

// AppPlugin defines a AppPlugin function signature.
type AppPlugin func(context.Context, *fiber.App) error

// Plugin defines a Plugin function signature.
type Plugin func(context.Context, *Options) (ConfigPlugin, AppPlugin)

// Server represents a fiber server.
type Server struct {
	app     *fiber.App
	options *Options
}

// NewServerWithConfigPath new fiber server with options from config path.
func NewServerWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*Server, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewServerWithOptions(ctx, opts, plugins...), nil
}

// NewServer new fiber server with default options.
func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, options, plugins...)
}

// NewServerWithOptions new fiber server with options.
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

// App returns fiber app instance.
func (s *Server) App() *fiber.App {
	return s.app
}

// Get registers handlers for method GET at path.
func (s *Server) Get(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Get(path, handlers...)
}

// Post registers handlers for method POST at path.
func (s *Server) Post(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Post(path, handlers...)
}

// Delete registers handlers for method DELETE at path.
func (s *Server) Delete(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Delete(path, handlers...)
}

// Head registers handlers for method HEAD at path.
func (s *Server) Head(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Head(path, handlers...)
}

// Put registers handlers for method PUT at path.
func (s *Server) Put(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Put(path, handlers...)
}

// Patch registers handlers for method PATCH at path.
func (s *Server) Patch(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Patch(path, handlers...)
}

// Add registers handlers for method at path
func (s *Server) Add(method string, path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Add(method, path, handlers...)
}

// Connect registers a route for CONNECT methods that establishes a tunnel to the
// server identified by the target resource.
func (s *Server) Connect(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Connect(path, handlers...)
}

// All registers handlers for all methods at path.
func (s *Server) All(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.All(path, handlers...)
}

// Group is used for Routes with common prefix to define a new sub-router with optional middleware.
//  api := app.Group("/api")
//  api.Get("/users", handler)
func (s *Server) Group(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Group(path, handlers...)
}

// Options registers handlers for OPTION method at path.
func (s *Server) Options(path string, handlers ...fiber.Handler) fiber.Router {
	return s.app.Options(path, handlers...)
}

// Serve starts app to serve at specified address.
func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)
	logger.Infof("starting fiber Server. https://gofiber.io/")

	addr := ":" + strconv.Itoa(s.options.Port)

	logger.Error(s.app.Listen(addr))
}

// Shutdown shut server down gracefully.
func (s *Server) Shutdown(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Error(s.app.Shutdown())
}
