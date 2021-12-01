package status

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	e "github.com/labstack/echo/v4"
)

// Register registers a new status plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewStatusWithOptions(o)
	return h.Register(ctx, server)
}

// Status represents status plugin for echo server.
type Status struct {
	options *Options
}

// NewStatusWithOptions returns a new status plugin with options.
func NewStatusWithOptions(options *Options) *Status {
	return &Status{options: options}
}

// NewStatusWithConfigPath returns a new status plugin with options from config path.
func NewStatusWithConfigPath(path string) (*Status, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStatusWithOptions(o), nil
}

// NewStatus returns a new status plugin with default options.
func NewStatus() *Status {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewStatusWithOptions(o)
}

// Register registers this status plugin for echo server.
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
