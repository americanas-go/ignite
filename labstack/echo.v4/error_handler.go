package echo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/americanas-go/errors"
	response "github.com/americanas-go/rest-response"
	"github.com/go-playground/validator/v10"
	e "github.com/labstack/echo/v4"
)

// ErrorHandlerString implements plain text content type error handler.
func ErrorHandlerString(err error, c e.Context) {
	errorHandler(err, c, e.MIMETextPlain)
}

// ErrorHandlerJSON implements JSON content type error handler.
func ErrorHandlerJSON(err error, c e.Context) {
	errorHandler(err, c, e.MIMEApplicationJSON)
}

func errorHandler(err error, c e.Context, contentType string) {
	var (
		status  int
		message string
	)
	if echoErr, ok := err.(*e.HTTPError); ok {
		status = echoErr.Code
		message = fmt.Sprintf("%v", echoErr.Message)
	} else {
		status = ErrorStatusCode(err)
		message = err.Error()
	}

	var er error
	if c.Request().Method == http.MethodHead {
		er = c.NoContent(status)
	} else {
		switch contentType {
		case e.MIMEApplicationJSON:
			er = c.JSON(status, response.Error{HttpStatusCode: status, ErrorCode: strconv.Itoa(status), Message: message})
		default:
			er = c.String(status, message)
		}
	}
	if er != nil {
		c.Logger().Error(er)
	}
}

// ErrorStatusCode translates to the respective status code.
func ErrorStatusCode(err error) int {

	switch {
	case errors.IsNotFound(err):
		return http.StatusNotFound
	case errors.IsMethodNotAllowed(err):
		return http.StatusMethodNotAllowed
	case errors.IsNotValid(err) || errors.IsBadRequest(err):
		return http.StatusBadRequest
	case errors.IsServiceUnavailable(err):
		return http.StatusServiceUnavailable
	case errors.IsConflict(err) || errors.IsAlreadyExists(err):
		return http.StatusConflict
	case errors.IsNotImplemented(err) || errors.IsNotProvisioned(err):
		return http.StatusNotImplemented
	case errors.IsUnauthorized(err):
		return http.StatusUnauthorized
	case errors.IsForbidden(err):
		return http.StatusForbidden

	default:
		if _, ok := err.(validator.ValidationErrors); ok {
			return http.StatusUnprocessableEntity
		}
		return http.StatusInternalServerError
	}
}
