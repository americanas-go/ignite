package chi

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/net/http/server"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5"
)

// Config represents chi plugin configuration.
type Config struct {
	Middlewares []func(http.Handler) http.Handler
	Handlers    []ConfigHandler
	Routes      []ConfigRouter
}

// ConfigHandler represents chi handler configuration.
type ConfigHandler struct {
	Handler http.Handler
	Pattern string
}

// ConfigRouter represents chi router configuration.
type ConfigRouter struct {
	Method      string
	HandlerFunc http.HandlerFunc
	Pattern     string
}

// Plugin defines a function to process plugin.
type Plugin func(context.Context) (*Config, error)

// Server represents a chi http server.
type Server struct {
	mux        *chi.Mux
	opts       *server.Options
	httpServer *http.Server
}

// NewServer creates a new http server for chi with plugins.
func NewServer(ctx context.Context, plugins ...Plugin) *Server {
	opt, err := server.NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(ctx, opt, plugins...)
}

// NewServer creates a new http server for chi with options from path and plugins.
func NewServerWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*Server, error) {
	o, err := server.NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewServerWithOptions(ctx, o, plugins...), nil
}

// NewServer creates a new http server for chi with options and plugins
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

// Mux returns the chi http multiplexer.
func (s *Server) Mux() *chi.Mux {
	return s.mux
}

// HtpServer returns the underlying http server.
func (s *Server) HttpServer() *http.Server {
	return s.httpServer
}

// Serve handles requests on incoming connections.
func (s *Server) Serve(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("started chi http Server [%s]", s.httpServer.Addr)
	logger.Error(s.httpServer.ListenAndServe())
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (s *Server) Shutdown(ctx context.Context) {
	logger := log.FromContext(ctx)
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	logger.Info("stopped chi http Server")
}

// Get adds the route `pattern` that matches a GET http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Get(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Get(pattern, handlerFn)
}

// Post adds the route `pattern` that matches a POST http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Post(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Post(pattern, handlerFn)
}

// Delete adds the route `pattern` that matches a DELETE http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Delete(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Delete(pattern, handlerFn)
}

// Head adds the route `pattern` that matches a HEAD http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Head(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Head(pattern, handlerFn)
}

// Put adds the route `pattern` that matches a PUT http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Put(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Put(pattern, handlerFn)
}

// Patch adds the route `pattern` that matches a PATCH http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Patch(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Patch(pattern, handlerFn)
}

// Connect adds the route `pattern` that matches a CONNECT http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Connect(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Connect(pattern, handlerFn)
}

// Group creates a new inline-Mux with a fresh middleware stack. It's useful for a group of handlers along the same routing path that use an additional set of middlewares. See _examples/.
func (s *Server) Group(fn func(r chi.Router)) {
	s.mux.Group(fn)
}

// Route creates a new Mux with a fresh middleware stack and mounts it along the `pattern` as a subrouter. Effectively, this is a short-hand call to Mount. See _examples/.
func (s *Server) Route(pattern string, fn func(r chi.Router)) {
	s.mux.Route(pattern, fn)
}

// Handle adds the route `pattern` that matches any http method to execute the `handler` http.Handler.
func (s *Server) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

// HandleFunc adds the route `pattern` that matches any http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	s.mux.HandleFunc(pattern, handlerFn)
}

// MethodFunc adds the route `pattern` that matches `method` http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) MethodFunc(method string, pattern string, handlerFn http.HandlerFunc) {
	s.mux.MethodFunc(method, pattern, handlerFn)
}

// Method adds the route `pattern` that matches `method` http method to execute the `handler` http.Handler.
func (s *Server) Method(method string, pattern string, handler http.Handler) {
	s.mux.Method(method, pattern, handler)
}

// Match searches the routing tree for a handler that matches the method/path. It's similar to routing a http request, but without executing the handler thereafter.
//
// Note: the *Context state is updated during execution, so manage the state carefully or make a NewRouteContext().
func (s *Server) Match(ctx *chi.Context, method string, path string) {
	s.mux.Match(ctx, method, path)
}

// Trace adds the route `pattern` that matches a TRACE http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Trace(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Trace(pattern, handlerFn)
}

// Use appends a middleware handler to the Mux middleware stack.
//
// The middleware stack for any Mux will execute before searching for a matching route to a specific handler, which provides opportunity to respond early, change the course of the request execution, or set request-scoped values for the next http.Handler.
func (s *Server) Use(middlewares ...func(http.Handler) http.Handler) {
	s.mux.Use(middlewares...)
}

// With adds inline middlewares for an endpoint handler.
func (s *Server) With(middlewares ...func(http.Handler) http.Handler) chi.Router {
	return s.mux.With(middlewares...)
}

// Options adds the route `pattern` that matches a OPTIONS http method to execute the `handlerFn` http.HandlerFunc.
func (s *Server) Options(pattern string, handlerFn http.HandlerFunc) {
	s.mux.Options(pattern, handlerFn)
}
