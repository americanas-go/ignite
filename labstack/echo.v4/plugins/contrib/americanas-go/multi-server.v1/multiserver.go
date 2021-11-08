package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
	e "github.com/labstack/echo/v4"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewMultiServerWithOptions(o)
	return h.Register(ctx, server)
}

type MultiServer struct {
	options *Options
}

func NewMultiServerWithOptions(options *Options) *MultiServer {
	return &MultiServer{options: options}
}

func NewMultiServerWithConfigPath(path string) (*MultiServer, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewMultiServerWithOptions(o), nil
}

func NewMultiServer() *MultiServer {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewMultiServerWithOptions(o)
}

func (i *MultiServer) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	checkRoute := i.options.Route

	logger.Tracef("configuring multi server check router on %s in echo", checkRoute)

	server.GET(checkRoute, handler)

	logger.Debugf("multi server check router configured on %s in echo", checkRoute)

	return nil
}

func handler(c e.Context) error {

	status := http.StatusOK
	msg := "OK"

	if err := multiserver.Check(c.Request().Context()); err != nil {
		status = http.StatusServiceUnavailable
		msg = "Service Unavailable"
	}

	return c.String(status, msg)
}
