package error_handler

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
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
