package recoverer

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
)

// Register registers recoverer middleware for chi.
func Register(ctx context.Context) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewRecovererWithOptions(o)
	return n.Register(ctx)
}

// Recoverer struct which represents a recoverer middleware for chi.
type Recoverer struct {
	options *Options
}

// NewRecovererWithConfigPath returns a recoverer middleware with options from config path.
func NewRecovererWithConfigPath(path string) (*Recoverer, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRecovererWithOptions(o), nil
}

// NewRecovererWithOptions returns a recoverer middleware with options.
func NewRecovererWithOptions(options *Options) *Recoverer {
	return &Recoverer{options: options}
}

// Register registers this recoverer middleware on a new chi config.
func (d *Recoverer) Register(ctx context.Context) (*chi.Config, error) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling recoverer middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.Recoverer,
		},
	}, nil
}
