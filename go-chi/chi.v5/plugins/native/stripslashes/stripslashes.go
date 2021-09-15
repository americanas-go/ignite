package stripslashes

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
	n := NewStripSlashesWithOptions(o)
	return n.Register(ctx)
}

type StripSlashes struct {
	options *Options
}

func NewStripSlashesWithConfigPath(path string) (*StripSlashes, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewStripSlashesWithOptions(o), nil
}

func NewStripSlashesWithOptions(options *Options) *StripSlashes {
	return &StripSlashes{options: options}
}

func (d *StripSlashes) Register(ctx context.Context) (*chi.Config, error) {

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
