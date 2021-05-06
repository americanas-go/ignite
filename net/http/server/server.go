package server

import (
	"net/http"
)

// NewServerWithOptions returns a pointer with new Server
func NewServer(handler http.Handler) *http.Server {
	opt, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return NewServerWithOptions(handler, opt)
}

// NewServerWithOptions returns a pointer with new Server
func NewServerWithOptions(handler http.Handler, options *Options) *http.Server {
	return &http.Server{
		Addr:              options.Addr,
		Handler:           handler,
		MaxHeaderBytes:    options.MaxHeaderBytes,
		ReadTimeout:       options.ReadTimeout,
		ReadHeaderTimeout: options.ReadHeaderTimeout,
		WriteTimeout:      options.WriteTimeout,
		IdleTimeout:       options.IdleTimeout,
	}
}
