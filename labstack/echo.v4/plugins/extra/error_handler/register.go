package error_handler

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	e "github.com/labstack/echo/v4"
)

// Register registers the error handler for echo.
//
// The error handler and converter middleware are registered.
// The handler handles how the application's errors will be sent in the response.
// And the converter converts application errors into echo errors.
//
// If the New Relic agent integration for Echo is used, this Register function must
// be called after the New Relic agent Register function.
func Register(ctx context.Context, server *echo.Server) error {
	if !IsEnabled() {
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
		status := echo.ErrorStatusCode(err)
		return e.NewHTTPError(status, err.Error())
	}
}
