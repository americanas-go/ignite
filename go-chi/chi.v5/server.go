package chi

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/net/http/server"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5"
)

type Config struct {
	Middlewares []func(http.Handler) http.Handler
	Handlers    []ConfigHandler
	Routes      []ConfigRouter
}

type ConfigHandler struct {
	Handler http.Handler
	Pattern string
}

type ConfigRouter struct {
	Method      string
	HandlerFunc http.HandlerFunc
	Pattern     string
}

type Plugin func(context.Context) (*Config, error)

type Server struct {
	mux        *chi.Mux
	opts       *server.Options
	httpServer *http.Server
}

func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	opt, err := server.NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, plugins...)
}

func NewServerWithOptions(ctx context.Context, opts *server.Options, plugins ...Plugin) *Server {

	mux := chi.NewRouter()

	var middlewares []func(http.Handler) http.Handler
	var handlers []ConfigHandler
	var routes []ConfigRouter

	for _, plugin := range plugins {

		var err error
		var config *Config

		if config, err = plugin(ctx); err != nil {
			panic(err)
		}

		if config != nil {

			if len(config.Middlewares) > 0 {
				middlewares = append(middlewares, config.Middlewares...)
			}

			if len(config.Handlers) > 0 {
				handlers = append(handlers, config.Handlers...)
			}

			if len(config.Routes) > 0 {
				routes = append(routes, config.Routes...)
			}

		}
	}

	if len(middlewares) > 0 {
		mux.Use(middlewares...)
	}

	if len(handlers) > 0 {
		for _, h := range handlers {
			mux.Handle(h.Pattern, h.Handler)
		}
	}

	if len(routes) > 0 {
		for _, r := range routes {
			mux.MethodFunc(r.Method, r.Pattern, r.HandlerFunc)
		}
	}

	httpServer := server.NewServerWithOptions(mux, opts)

	return &Server{mux: mux, httpServer: httpServer, opts: opts}
}

func (s *Server) Mux() *chi.Mux {
	return s.mux
}

func (s *Server) HttpServer() *http.Server {
	return s.httpServer
}

func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)

	logger.Infof("started chi http Server [%s]", s.httpServer.Addr)
	logger.Error(s.httpServer.ListenAndServe())
}

func (s *Server) Shutdown(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Error(s.httpServer.Shutdown(ctx))
}

func (s *Server) Get(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Get(pattern, handlerFn)
}

func (s *Server) Post(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Post(pattern, handlerFn)
}

func (s *Server) Delete(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Delete(pattern, handlerFn)
}

func (s *Server) Head(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Head(pattern, handlerFn)
}

func (s *Server) Put(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Put(pattern, handlerFn)
}

func (s *Server) Patch(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Patch(pattern, handlerFn)
}

func (s *Server) Connect(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Connect(pattern, handlerFn)
}

func (s *Server) Group(fn func(r chi.Router)) {
	s.mux.Group(fn)
}

func (s *Server) Route(pattern string, fn func(r chi.Router)) {
	s.mux.Route(pattern, fn)
}

func (s *Server) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *Server) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	s.mux.HandleFunc(pattern, handlerFn)
}

func (s *Server) MethodFunc(method string, pattern string, handlerFn http.HandlerFunc) {
	s.mux.MethodFunc(method, pattern, handlerFn)
}

func (s *Server) Method(method string, pattern string, handler http.Handler) {
	s.mux.Method(method, pattern, handler)
}

func (s *Server) Match(ctx *chi.Context, method string, path string) {
	s.mux.Match(ctx, method, path)
}

func (s *Server) Trace(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Trace(pattern, handlerFn)
}

func (s *Server) Use(middlewares ...func(http.Handler) http.Handler) {
	s.mux.Use(middlewares...)
}

func (s *Server) With(middlewares ...func(http.Handler) http.Handler) chi.Router {
	return s.mux.With(middlewares...)
}

func (s *Server) Options(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Options(pattern, handlerFn)
}
