package requestid

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewRequestIDWithOptions(o)
	return h.Register(ctx, server)
}

type RequestID struct {
	options *Options
}

func NewRequestIDWithOptions(options *Options) *RequestID {
	return &RequestID{options: options}
}

func NewRequestIDWithConfigPath(path string) (*RequestID, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRequestIDWithOptions(o), nil
}

func NewRequestID() *RequestID {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewRequestIDWithOptions(o)
}

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
