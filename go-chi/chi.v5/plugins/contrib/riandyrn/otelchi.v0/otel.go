package otel

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	c "github.com/go-chi/chi/v5"
	"github.com/riandyrn/otelchi"
)

type OTel struct {
	options *Options
}

func NewOTelWithConfigPath(path string) (*OTel, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOTelWithOptions(o), nil
}

func NewOTelWithOptions(options *Options) *OTel {
	return &OTel{options: options}
}

func (d *OTel) Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {

	if !d.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling otel middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			otelchi.Middleware(d.options.ServerName, otelchi.WithChiRoutes(mux)),
		},
	}, nil

}

func Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	n := NewOTelWithOptions(o)
	return n.Register(ctx, mux)
}
