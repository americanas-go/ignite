package server

import (
	"context"
	"fmt"
	"net"

	"github.com/americanas-go/log"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

// Server represents a drpc server.
type Server struct {
	server  *drpcserver.Server
	mux     *drpcmux.Mux
	options *Options
}

// NewServer returns a new drpc server with default options.
func NewServer(ctx context.Context) (*Server, error) {
	opt, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewServerWithOptions(ctx, opt)
}

// NewServerWithConfigPath returns a new drpc server with options from config path.
func NewServerWithConfigPath(ctx context.Context, path string) (*Server, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewServerWithOptions(ctx, options)
}

// NewServerWithOptions returns a new drpc server with options.
func NewServerWithOptions(ctx context.Context, opt *Options) (*Server, error) {

	// create a drpc RPC mux
	m := drpcmux.New()

	// create a drpc server
	server := drpcserver.New(m)

	return &Server{
		mux:     m,
		server:  server,
		options: opt,
	}, nil
}

// Server returns the wrapped drpc server instance.
func (s *Server) Server() *drpcserver.Server {
	return s.server
}

// Mux returns the wrapped drpc mux instance.
func (s *Server) Mux() *drpcmux.Mux {
	return s.mux
}

// Serve starts drpc server.
func (s *Server) Serve(ctx context.Context) {

	logger := log.FromContext(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.options.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err.Error())
	}

	logger.Infof("drpc server started on port %v", s.options.Port)

	logger.Error(s.server.Serve(ctx, lis))
}

// Shutdown stops drpc server gracefully.
func (s *Server) Shutdown(ctx context.Context) {}
