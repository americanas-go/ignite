package stripslashes

import (
	"context"
	c "github.com/go-chi/chi/v5"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
)

// Register registers stripslashes plugin for chi.
func Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewStripSlashesWithOptions(o)
	return n.Register(ctx, mux)
}

// StripSlashes struct which represents a stripslashes plugin from chi.
type StripSlashes struct {
	options *Options
}

// NewStripSlashesWithConfigPath returns a new stripslashes with options from config path.
func NewStripSlashesWithConfigPath(path string) (*StripSlashes, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStripSlashesWithOptions(o), nil
}

// NewStripSlashesWithOptions returns a new stripslashes plugin with options.
func NewStripSlashesWithOptions(options *Options) *StripSlashes {
	return &StripSlashes{options: options}
}

// Register registers this stripslashes plugin for a new chi config.
func (d *StripSlashes) Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling stripSlashes middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.StripSlashes,
		},
	}, nil
}
