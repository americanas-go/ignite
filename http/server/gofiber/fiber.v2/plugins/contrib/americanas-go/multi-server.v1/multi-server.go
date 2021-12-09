package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
	f "github.com/gofiber/fiber/v2"
)

// Register registers a new multiserver plugin for fiber.
func Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	l := NewMultiServer()
	return l.Register(ctx, options)
}

// MultiServer represents multiserver plugin for fiber.
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

// Register registers this multiserver plugin for fiber.
func (i *MultiServer) Register(ctx context.Context, options *fiber.Options) (fiber.ConfigPlugin, fiber.AppPlugin) {
	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	checkRoute := i.options.Route

	logger.Tracef("configuring multi server check router on %s in fiber", checkRoute)

	return nil, func(ctx context.Context, app *f.App) error {

		app.Get(checkRoute, handler)

		logger.Debugf("multi server check router configured on %s in fiber", checkRoute)
		return nil
	}

}

func handler(c *f.Ctx) error {

	status := http.StatusOK
	msg := "OK"

	if err := multiserver.Check(c.Context()); err != nil {
		status = http.StatusServiceUnavailable
		msg = "Service Unavailable"
	}

	return c.Status(status).SendString(msg)
}
