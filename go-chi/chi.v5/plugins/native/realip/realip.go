package realip

import (
	"context"
	c "github.com/go-chi/chi/v5"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
)

// Register registers realIP middleware for chi.
func Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewRealIPWithOptions(o)
	return n.Register(ctx, mux)
}

// RealIP struct which represents a realIP plugin
type RealIP struct {
	options *Options
}

// NewRealIPWithConfigPath returns realIP plugin with options from config path.
func NewRealIPWithConfigPath(path string) (*RealIP, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRealIPWithOptions(o), nil
}

// NewRealIPWithOptions returns realIP plugin with options.
func NewRealIPWithOptions(options *Options) *RealIP {
	return &RealIP{options: options}
}

// Register registers the realIP plugin on a new chi config.
func (d *RealIP) Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
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
