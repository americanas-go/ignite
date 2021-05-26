package error_handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	"github.com/go-playground/validator/v10"
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
				errorHandlerString(err, c)
			} else {
				errorHandlerJSON(err, c)
			}
		}

	}
}

func errorHandlerString(err error, c e.Context) {

	var er error
	var status int
	message := err.Error()

	if errors.IsNotFound(err) {
		status = http.StatusNotFound
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		status = http.StatusBadRequest
	} else if errors.IsServiceUnavailable(err) {
		status = http.StatusServiceUnavailable
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		status = http.StatusConflict
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		status = http.StatusNotImplemented
	} else if errors.IsUnauthorized(err) {
		status = http.StatusUnauthorized
	} else if errors.IsForbidden(err) {
		status = http.StatusForbidden
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			status = http.StatusUnprocessableEntity
		case *e.HTTPError:
			status = t.Code
			message = fmt.Sprintf("%v", t.Message)
		default:
			status = http.StatusInternalServerError
		}
	}

	if c.Request().Method == http.MethodHead {
		er = c.NoContent(status)
	} else {
		er = c.String(status, message)
	}

	if er != nil {
		c.Logger().Error(er)
	}
}

func errorHandlerJSON(err error, c e.Context) {

	var er error
	var status int
	message := err.Error()

	if errors.IsNotFound(err) {
		status = http.StatusNotFound
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		status = http.StatusBadRequest
	} else if errors.IsServiceUnavailable(err) {
		status = http.StatusServiceUnavailable
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		status = http.StatusConflict
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		status = http.StatusNotImplemented
	} else if errors.IsUnauthorized(err) {
		status = http.StatusUnauthorized
	} else if errors.IsForbidden(err) {
		status = http.StatusForbidden
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			status = http.StatusUnprocessableEntity
		case *e.HTTPError:
			status = t.Code
			message = fmt.Sprintf("%v", t.Message)
		default:
			status = http.StatusInternalServerError
		}
	}

	if c.Request().Method == http.MethodHead {
		er = c.NoContent(status)
	} else {
		er = c.JSON(status,
			response.Error{HttpStatusCode: status, Message: message})
	}

	if er != nil {
		c.Logger().Error(er)
	}

}
