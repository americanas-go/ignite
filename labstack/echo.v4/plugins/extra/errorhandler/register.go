package errorhandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *e.Echo) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("configuring error handler in echo")

	instance.HTTPErrorHandler = customHTTPErrorHandler

	logger.Debug("error handler successfully configured in echo")

	return nil
}

func customHTTPErrorHandler(err error, c e.Context) {
	code := http.StatusInternalServerError
	var msg interface{}
	if he, ok := err.(*e.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else {
		msg = err.Error()
	}

	resp := response.Error{HttpStatusCode: code, Message: fmt.Sprintf("%v", msg)}
	if err := echo.JSON(c, code, resp, nil); err != nil {
		c.Logger().Error(err)
	}
}
