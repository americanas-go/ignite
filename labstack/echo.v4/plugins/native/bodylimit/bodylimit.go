package bodylimit

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
	h := NewBodyLimitWithOptions(o)
	return h.Register(ctx, server)
}

type BodyLimit struct {
	options *Options
}

func NewBodyLimitWithOptions(options *Options) *BodyLimit {
	return &BodyLimit{options: options}
}

func NewBodyLimitWithConfigPath(path string) (*BodyLimit, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewBodyLimitWithOptions(o), nil
}

func NewBodyLimit() *BodyLimit {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewBodyLimitWithOptions(o)
}

func (i *BodyLimit) Register(ctx context.Context, server *echo.Server) error {
	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling body limit middleware in echo")

	server.Use(middleware.BodyLimit(i.options.Size))

	logger.Debug("body limit middleware successfully enabled in echo")

	return nil
}
