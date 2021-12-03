package cors

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4/middleware"
)

// Register registers a new CORS plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewCorsWithOptions(o)
	return h.Register(ctx, server)
}

// Cors represents CORS plugin for echo server.
type Cors struct {
	options *Options
}

// NewCorsWithOptions returns a new CORS plugin with options.
func NewCorsWithOptions(options *Options) *Cors {
	return &Cors{options: options}
}

// NewCorsWithConfigPath returns a new CORS plugin with options from config path.
func NewCorsWithConfigPath(path string) (*Cors, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCorsWithOptions(o), nil
}

// NewCors returns a new CORS plugin with default options.
func NewCors() *Cors {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewCorsWithOptions(o)
}

// Register registers this CORS plugin for echo server.
func (i *Cors) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling cors middleware in echo")

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     i.options.Allowed.Origins,
		AllowMethods:     i.options.Allowed.Methods,
		AllowHeaders:     i.options.Allowed.Headers,
		AllowCredentials: i.options.Allowed.Credentials,
		ExposeHeaders:    i.options.Exposed.Headers,
		MaxAge:           i.options.MaxAge,
	}))

	logger.Debug("cors middleware successfully enabled in echo")

	return nil
}
