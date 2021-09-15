package error_handler

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewErrorHandlerWithOptions(o)
	return h.Register(ctx, server)
}

type ErrorHandler struct {
	options *Options
}

func NewErrorHandlerWithOptions(options *Options) *ErrorHandler {
	return &ErrorHandler{options: options}
}

func NewErrorHandlerWithConfigPath(path string) (*ErrorHandler, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewErrorHandlerWithOptions(o), nil
}

func NewErrorHandler() *ErrorHandler {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewErrorHandlerWithOptions(o)
}

func (i *ErrorHandler) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("configuring error handler in echo")

	server.Instance().HTTPErrorHandler = errorHandler(server)
	server.Instance().Use(errorConverter)

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

func errorConverter(next e.HandlerFunc) e.HandlerFunc {
	return func(c e.Context) error {
		err := next(c)
		if err != nil {
			status := echo.ErrorStatusCode(err)
			return e.NewHTTPError(status, err.Error())
		}
		return nil
	}
}
