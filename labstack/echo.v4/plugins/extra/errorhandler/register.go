package errorhandler

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

	if server.Options().Type != "REST" {
		server.Instance().HTTPErrorHandler = errorHandler
	} else {
		server.Instance().HTTPErrorHandler = errorHandlerJSON
	}

	logger.Debug("error handler successfully configured in echo")

	return nil
}

func errorHandler(err error, c e.Context) {

	var er error

	if errors.IsNotFound(err) {
		er = c.String(http.StatusNotFound,
			err.Error())
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		er = c.String(http.StatusBadRequest,
			err.Error())
	} else if errors.IsServiceUnavailable(err) {
		er = c.String(http.StatusServiceUnavailable,
			err.Error())
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		er = c.String(http.StatusConflict,
			err.Error())
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		er = c.String(http.StatusNotImplemented,
			err.Error())
	} else if errors.IsUnauthorized(err) {
		er = c.String(http.StatusUnauthorized,
			err.Error())
	} else if errors.IsForbidden(err) {
		er = c.String(http.StatusForbidden,
			err.Error())
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			er = c.String(http.StatusUnprocessableEntity,
				t.Error())
		case *e.HTTPError:
			er = c.String(
				t.Code,
				t.Error())
		default:
			er = c.String(http.StatusInternalServerError,
				t.Error())
		}
	}

	if er != nil {
		c.Logger().Error(er)
	}
}

func errorHandlerJSON(err error, c e.Context) {

	var er error

	if errors.IsNotFound(err) {
		er = c.JSON(http.StatusNotFound,
			response.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()})
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		er = c.JSON(http.StatusBadRequest,
			response.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()})
	} else if errors.IsServiceUnavailable(err) {
		er = c.JSON(http.StatusServiceUnavailable,
			response.Error{HttpStatusCode: http.StatusServiceUnavailable, Message: err.Error()})
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		er = c.JSON(http.StatusConflict,
			response.Error{HttpStatusCode: http.StatusConflict, Message: err.Error()})
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		er = c.JSON(http.StatusNotImplemented,
			response.Error{HttpStatusCode: http.StatusNotImplemented, Message: err.Error()})
	} else if errors.IsUnauthorized(err) {
		er = c.JSON(http.StatusUnauthorized,
			response.Error{HttpStatusCode: http.StatusUnauthorized, Message: err.Error()})
	} else if errors.IsForbidden(err) {
		er = c.JSON(http.StatusForbidden,
			response.Error{HttpStatusCode: http.StatusForbidden, Message: err.Error()})
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			er = c.JSON(http.StatusUnprocessableEntity,
				response.NewUnprocessableEntity(t))
		case *e.HTTPError:
			er = c.JSON(
				t.Code,
				response.Error{HttpStatusCode: t.Code, Message: fmt.Sprintf("%v", t.Message)})
		default:
			er = c.JSON(http.StatusInternalServerError,
				response.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()})
		}
	}

	if er != nil {
		c.Logger().Error(er)
	}

}
