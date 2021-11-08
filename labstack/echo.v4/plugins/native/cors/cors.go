package cors

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewCorsWithOptions(o)
	return h.Register(ctx, server)
}

type Cors struct {
	options *Options
}

func NewCorsWithOptions(options *Options) *Cors {
	return &Cors{options: options}
}

func NewCorsWithConfigPath(path string) (*Cors, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCorsWithOptions(o), nil
}

func NewCors() *Cors {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewCorsWithOptions(o)
}

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
