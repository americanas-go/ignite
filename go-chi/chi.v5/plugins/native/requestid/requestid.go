package requestid

import (
	"context"
	c "github.com/go-chi/chi/v5"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/chi/v5/middleware"
)

// Register registers requestID middleware for chi.
func Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewRequestIDWithOptions(o)
	return n.Register(ctx, mux)
}

// RequestID struct which represents a requestID middleware for chi.
type RequestID struct {
	options *Options
}

// NewRequestIDWithConfigPath returns a new requestID plugin with options from config path.
func NewRequestIDWithConfigPath(path string) (*RequestID, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRequestIDWithOptions(o), nil
}

// NewRequestIDWithOptions returns a new requestID plugin with options.
func NewRequestIDWithOptions(options *Options) *RequestID {
	return &RequestID{options: options}
}

// Register registers this requestID plugin to a new chi config.
func (d *RequestID) Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling requestID middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.RequestID,
		},
	}, nil
}
