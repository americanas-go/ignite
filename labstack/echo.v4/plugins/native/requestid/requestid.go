package requestid

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4/middleware"
)

// Register registers a new requestID plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewRequestIDWithOptions(o)
	return h.Register(ctx, server)
}

// RequestID represents requestID plugin for echo server.
type RequestID struct {
	options *Options
}

// NewRequestIDWithOptions returns a new requestID plugin with options.
func NewRequestIDWithOptions(options *Options) *RequestID {
	return &RequestID{options: options}
}

// NewRequestIDWithConfigPath returns a new requestID plugin with options from config path.
func NewRequestIDWithConfigPath(path string) (*RequestID, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRequestIDWithOptions(o), nil
}

// NewRequestID returns a new requestID plugin with default options.
func NewRequestID() *RequestID {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewRequestIDWithOptions(o)
}

// Register registers this requestID plugin for echo server.
func (i *RequestID) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling requestID middleware in echo")

	server.Use(middleware.RequestID())

	logger.Debug("requestID middleware successfully enabled in echo")

	return nil
}
