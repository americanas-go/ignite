package health

import (
	"context"

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
	h := NewHealthWithOptions(o)
	return h.Register(ctx, server)
}

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

func (i *Health) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	healthRoute := i.options.Route

	logger.Tracef("configuring health router on %s in echo", healthRoute)

	server.GET(healthRoute, handler)

	logger.Debugf("health router configured on %s in echo", healthRoute)

	return nil
}

func handler(c e.Context) error {

	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	resp, httpCode := response.NewHealth(ctx)

	return c.JSON(httpCode, resp)
}
