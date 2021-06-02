package echo

import (
	"fmt"
	"net/http"

	"github.com/americanas-go/errors"
	response "github.com/americanas-go/rest-response"
	"github.com/go-playground/validator/v10"
	e "github.com/labstack/echo/v4"
)

func ErrorHandlerString(err error, c e.Context) {

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

func ErrorHandlerJSON(err error, c e.Context) {

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
