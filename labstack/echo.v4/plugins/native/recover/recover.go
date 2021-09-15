package gzip

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
	h := NewRecoverWithOptions(o)
	return h.Register(ctx, server)
}

type Recover struct {
	options *Options
}

func NewRecoverWithOptions(options *Options) *Recover {
	return &Recover{options: options}
}

func NewRecoverWithConfigPath(path string) (*Recover, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRecoverWithOptions(o), nil
}

func NewRecover() *Recover {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewRecoverWithOptions(o)
}

func (i *Recover) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling recover middleware in echo")

	server.Use(middleware.Recover())

	logger.Debug("recover middleware successfully enabled in echo")

	return nil
}
