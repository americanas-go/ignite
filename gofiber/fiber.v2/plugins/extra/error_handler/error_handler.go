package error_handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	"github.com/go-playground/validator/v10"
	f "github.com/gofiber/fiber/v2"
)

// Register registers a new error handler plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	n := NewErrorHandlerWithOptions(o)
	return n.Register(ctx, options)
}

// ErrorHandler represents a new error handler plugin for fiber.
type ErrorHandler struct {
	options *Options
}

// NewErrorHandlerWithConfigPath returns a new error handler plugin with options from config path.
func NewErrorHandlerWithConfigPath(path string) (*ErrorHandler, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewErrorHandlerWithOptions(o), nil
}

// NewErrorHandlerWithOptions returns a new error handler plugin with options.
func NewErrorHandlerWithOptions(options *Options) *ErrorHandler {
	return &ErrorHandler{options: options}
}

// Register registers this error handler plugin for fiber.
func (d *ErrorHandler) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("setting error handler in fiber")

	return func(ctx context.Context, config *f.Config) error {

		if options.Type != "REST" {
			config.ErrorHandler = d.errorHandler
		} else {
			config.ErrorHandler = d.errorHandlerJSON
		}

		logger.Debug("error handler successfully sets in fiber")

		return nil
	}, nil
}

func (d *ErrorHandler) print(ctx context.Context, err error, tp int) {
	logger := log.FromContext(ctx)
	if tp == 5 && !d.options.Logger.Print5xx {
		return
	} else if tp == 4 && !d.options.Logger.Print4xx {
		return
	} else {
		msg := err.Error()
		if d.options.Logger.PrintStackTrace {
			msg = errors.ErrorStack(err)
		}
		l := logger.FromContext(ctx).WithError(err)
		switch tp {
		case 5:
			l.Errorf(msg)
		default:
			l.Warnf(msg)
		}
	}
}

func (d *ErrorHandler) errorHandler(c *f.Ctx, err error) error {
	if errors.IsNotFound(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusNotFound).SendString(err.Error())
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	} else if errors.IsServiceUnavailable(err) {
		d.print(c.UserContext(), err, 5)
		return c.Status(http.StatusServiceUnavailable).SendString(err.Error())
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusConflict).SendString(err.Error())
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusNotImplemented).SendString(err.Error())
	} else if errors.IsUnauthorized(err) {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	} else if errors.IsForbidden(err) {
		return c.Status(http.StatusForbidden).SendString(err.Error())
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			d.print(c.UserContext(), err, 4)
			return c.Status(http.StatusUnprocessableEntity).SendString("The server understands the content type " +
				"of the request entity but was unable to process the contained instructions.")
		case *f.Error:
			d.print(c.UserContext(), err, 4)
			return c.Status(t.Code).SendString(t.Message)
		default:
			d.print(c.UserContext(), err, 5)
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
	}
}

func (d *ErrorHandler) errorHandlerJSON(c *f.Ctx, err error) error {

	if errors.IsNotFound(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusNotFound).JSON(
			response.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()})
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusBadRequest).JSON(
			response.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()})
	} else if errors.IsServiceUnavailable(err) {
		d.print(c.UserContext(), err, 5)
		return c.Status(http.StatusServiceUnavailable).JSON(
			response.Error{HttpStatusCode: http.StatusServiceUnavailable, Message: err.Error()})
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusConflict).JSON(
			response.Error{HttpStatusCode: http.StatusConflict, Message: err.Error()})
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusNotImplemented).JSON(
			response.Error{HttpStatusCode: http.StatusNotImplemented, Message: err.Error()})
	} else if errors.IsUnauthorized(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusUnauthorized).JSON(
			response.Error{HttpStatusCode: http.StatusUnauthorized, Message: err.Error()})
	} else if errors.IsForbidden(err) {
		d.print(c.UserContext(), err, 4)
		return c.Status(http.StatusForbidden).JSON(
			response.Error{HttpStatusCode: http.StatusForbidden, Message: err.Error()})
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			d.print(c.UserContext(), err, 4)
			return c.Status(http.StatusUnprocessableEntity).JSON(
				response.NewUnprocessableEntity(t))
		case *f.Error:
			fs := strconv.Itoa(t.Code)[0:1]
			if tp, err := strconv.Atoi(fs); err != nil {
				d.print(c.UserContext(), err, tp)
			}
			return c.Status(t.Code).JSON(
				response.Error{HttpStatusCode: t.Code, Message: t.Message})
		default:
			d.print(c.UserContext(), err, 5)
			return c.Status(http.StatusInternalServerError).JSON(
				response.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()})
		}
	}

}
