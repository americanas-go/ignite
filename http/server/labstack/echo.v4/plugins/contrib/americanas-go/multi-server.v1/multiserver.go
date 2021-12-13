package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/http/server/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
	e "github.com/labstack/echo/v4"
)

// Register registers a new multiserver plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewMultiServerWithOptions(o)
	return h.Register(ctx, server)
}

// MultiServer represents multiserver plugin for echo server.
type MultiServer struct {
	options *Options
}

// NewMultiServerWithOptions returns a new multiserver plugin with options.
func NewMultiServerWithOptions(options *Options) *MultiServer {
	return &MultiServer{options: options}
}

// NewMultiServerWithConfigPath returns a new multiserver plugin with options from config path.
func NewMultiServerWithConfigPath(path string) (*MultiServer, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewMultiServerWithOptions(o), nil
}

// NewMultiServer returns a new multiserver plugin with default options.
func NewMultiServer() *MultiServer {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewMultiServerWithOptions(o)
}

// Register registers this multiserver plugin for echo server.
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
