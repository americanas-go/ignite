package cors

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/go-chi/cors"
)

// Register registers cors middleware for chi.
func Register(ctx context.Context) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewCorsWithOptions(o)
	return n.Register(ctx)
}

type Cors struct {
	options *Options
}

func NewCorsWithConfigPath(path string) (*Cors, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewCorsWithOptions(o), nil
}

func NewCorsWithOptions(options *Options) *Cors {
	return &Cors{options: options}
}

func (d *Cors) Register(ctx context.Context) (*chi.Config, error) {

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling cors middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			cors.Handler(cors.Options{
				AllowedOrigins:   d.options.Allowed.Origins,
				AllowedMethods:   d.options.Allowed.Methods,
				AllowedHeaders:   d.options.Allowed.Headers,
				AllowCredentials: d.options.Allowed.Credentials,
				ExposedHeaders:   d.options.Exposed.Headers,
				MaxAge:           d.options.MaxAge,
			}),
		},
	}, nil
}
