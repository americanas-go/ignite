package bodylimit

import (
	"context"

	"github.com/americanas-go/ignite/labstack/echo.v4"
	"github.com/americanas-go/log"
	"github.com/labstack/echo/v4/middleware"
)

// Register registers a new bodylimit plugin for echo server.
func Register(ctx context.Context, server *echo.Server) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewBodyLimitWithOptions(o)
	return h.Register(ctx, server)
}

// BodyLimit represents bodylimit plugin for echo server.
type BodyLimit struct {
	options *Options
}

// NewBodyLimitWithOptions returns a new bodylimit plugin with options.
func NewBodyLimitWithOptions(options *Options) *BodyLimit {
	return &BodyLimit{options: options}
}

// NewBodyLimitWithConfigPath returns a new bodylimit plugin with options from config path.
func NewBodyLimitWithConfigPath(path string) (*BodyLimit, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewBodyLimitWithOptions(o), nil
}

// NewBodyLimit returns a new bodylimit plugin with default options.
func NewBodyLimit() *BodyLimit {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewBodyLimitWithOptions(o)
}

// Register registers this bodylimit plugin for echo server.
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
