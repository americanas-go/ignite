package realip

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(ctx context.Context) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewRealIPWithOptions(o)
	return n.Register(ctx)
}

type RealIP struct {
	options *Options
}

func NewRealIPWithConfigPath(path string) (*RealIP, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRealIPWithOptions(o), nil
}

func NewRealIPWithOptions(options *Options) *RealIP {
	return &RealIP{options: options}
}

func (d *RealIP) Register(ctx context.Context) (*chi.Config, error) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling realIP middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.RealIP,
		},
	}, nil
}
