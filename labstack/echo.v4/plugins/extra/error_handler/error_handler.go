package error_handler

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
)

// Register registers a new error handler plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewErrorHandlerWithOptions(o)
	return h.Register(ctx, server)
}

// ErrorHandler represents error handler plugin for echo server.
type ErrorHandler struct {
	options *Options
}

// NewErrorHandlerWithOptions returns a new error handler plugin with options.
func NewErrorHandlerWithOptions(options *Options) *ErrorHandler {
	return &ErrorHandler{options: options}
}

// NewErrorHandlerWithConfigPath returns a new error handler plugin with options from config path.
func NewErrorHandlerWithConfigPath(path string) (*ErrorHandler, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewErrorHandlerWithOptions(o), nil
}

// NewErrorHandler returns a new error handler plugin with default options.
func NewErrorHandler() *ErrorHandler {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewErrorHandlerWithOptions(o)
}

// Register registers this error handler plugin for echo server.
func (i *ErrorHandler) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("configuring error handler in echo")

	server.Instance().HTTPErrorHandler = errorHandler(server)

	logger.Debug("error handler successfully configured in echo")

	return nil
}

func errorHandler(server *echo.Server) func(err error, c e.Context) {
	return func(err error, c e.Context) {

		if !c.Response().Committed {
			if server.Options().Type != "REST" {
				echo.ErrorHandlerString(err, c)
			} else {
				echo.ErrorHandlerJSON(err, c)
			}
		}

	}
}
