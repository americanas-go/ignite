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
	h := NewGzipWithOptions(o)
	return h.Register(ctx, server)
}

type Gzip struct {
	options *Options
}

func NewGzipWithOptions(options *Options) *Gzip {
	return &Gzip{options: options}
}

func NewGzipWithConfigPath(path string) (*Gzip, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewGzipWithOptions(o), nil
}

func NewGzip() *Gzip {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewGzipWithOptions(o)
}

func (i *Gzip) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling gzip middleware in echo")

	server.Use(middleware.Gzip())

	logger.Debug("gzip middleware successfully enabled in echo")

	return nil
}