package status

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewStatusWithOptions(o)
	return h.Register(ctx, server)
}

type Status struct {
	options *Options
}

func NewStatusWithOptions(options *Options) *Status {
	return &Status{options: options}
}

func NewStatusWithConfigPath(path string) (*Status, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStatusWithOptions(o), nil
}

func NewStatus() *Status {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewStatusWithOptions(o)
}

func (i *Status) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	statusRoute := i.options.Route

	logger.Tracef("configuring status router on %s in echo", statusRoute)

	server.GET(statusRoute, handler)

	logger.Debugf("status router configured on %s in echo", statusRoute)

	return nil
}

func handler(c e.Context) error {
	return c.JSON(http.StatusOK, response.NewResourceStatus())
}
