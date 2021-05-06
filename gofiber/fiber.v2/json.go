package fiber

import (
	"net/http"

	"github.com/americanas-go/errors"
	response "github.com/americanas-go/rest-response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func JSON(c *fiber.Ctx, code int, i interface{}, err error) error {

	if err != nil {
		return JSONError(c, err)
	}

	if i == nil {
		c.Status(http.StatusNoContent)
		return nil
	}

	return c.Status(code).JSON(i)
}

func JSONError(c *fiber.Ctx, err error) error {

	if errors.IsNotFound(err) {
		return c.Status(http.StatusNotFound).JSON(
			response.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()})
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		return c.Status(http.StatusBadRequest).JSON(
			response.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()})
	} else if errors.IsServiceUnavailable(err) {
		return c.Status(http.StatusServiceUnavailable).JSON(
			response.Error{HttpStatusCode: http.StatusServiceUnavailable, Message: err.Error()})
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			return c.Status(http.StatusUnprocessableEntity).JSON(
				response.NewUnprocessableEntity(t))
		default:
			return c.Status(http.StatusInternalServerError).JSON(
				response.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()})
		}
	}

}
