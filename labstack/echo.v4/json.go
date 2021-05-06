package echo

import (
	"net/http"

	"github.com/americanas-go/errors"
	response "github.com/americanas-go/rest-response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func JSON(c echo.Context, code int, i interface{}, err error) error {

	if err != nil {

		return JSONError(c, err)

	}

	if i == nil {
		return c.NoContent(code)
	}

	return json(c, code, i)
}

func JSONError(c echo.Context, err error) error {

	if errors.IsNotFound(err) {
		return json(c,
			http.StatusNotFound,
			response.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()})
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		return json(c,
			http.StatusBadRequest,
			response.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()})
	} else if errors.IsServiceUnavailable(err) {
		return json(c,
			http.StatusServiceUnavailable,
			response.Error{HttpStatusCode: http.StatusServiceUnavailable, Message: err.Error()})
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			return json(c,
				http.StatusUnprocessableEntity,
				response.NewUnprocessableEntity(t))
		default:
			return json(c,
				http.StatusInternalServerError,
				response.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()})
		}
	}

}

func json(c echo.Context, code int, response interface{}) error {

	if GetJSONPrettyEnabled() {
		return c.JSONPretty(code, response, "  ")
	}

	return c.JSON(code, response)
}
