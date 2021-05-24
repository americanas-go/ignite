package echo

import (
	"context"
	"strconv"

	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4"
)

type Plugin func(context.Context, *Server) error

type Server struct {
	instance *echo.Echo
	options  *Options
}

func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	opt, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, plugins...)
}

func NewServerWithOptions(ctx context.Context, opt *Options, plugins ...Plugin) *Server {

	instance := echo.New()

	instance.HideBanner = opt.HideBanner
	instance.Logger = WrapLogger(log.GetLogger())

	srv := &Server{instance: instance, options: opt}

	for _, plugin := range plugins {
		if err := plugin(ctx, srv); err != nil {
			panic(err)
		}
	}

	return srv
}

func (s *Server) Instance() *echo.Echo {
	return s.instance
}

func (s *Server) Options() *Options {
	return s.options
}
func (s *Server) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.GET(path, h, m...)
}

func (s *Server) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.POST(path, h, m...)
}

func (s *Server) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.PUT(path, h, m...)
}

func (s *Server) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.DELETE(path, h, m...)
}

func (s *Server) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.OPTIONS(path, h, m...)
}

func (s *Server) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.PATCH(path, h, m...)
}

func (s *Server) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.HEAD(path, h, m...)
}

func (s *Server) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.CONNECT(path, h, m...)
}

func (s *Server) Group(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return s.instance.Group(prefix, m...)
}

func (s *Server) Add(method string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.Add(method, path, h, m...)
}

func (s *Server) Any(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return s.instance.Any(path, h, m...)
}

func (s *Server) Use(middleware ...echo.MiddlewareFunc) {
	s.instance.Use(middleware...)
}

func (s *Server) Static(prefix, root string) *echo.Route {
	return s.instance.Static(prefix, root)
}

func (s *Server) Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route {
	return s.instance.Match(methods, path, handler, middleware...)
}

func (s *Server) Pre(middleware ...echo.MiddlewareFunc) {
	s.instance.Pre(middleware...)
}

func (s *Server) File(path, file string, m ...echo.MiddlewareFunc) *echo.Route {
	return s.instance.File(path, file, m...)
}

func (s *Server) Serve(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("starting echo Server. https://echo.labstack.com/")
	address := ":" + strconv.Itoa(s.options.Port)
	logger.Error(s.instance.Start(address))
}

func (s *Server) Shutdown(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Error(s.instance.Shutdown(ctx))
}
