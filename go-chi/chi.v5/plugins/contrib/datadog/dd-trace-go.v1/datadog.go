package datadog

import (
	"context"
	"net/http"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	c "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi.v5"
)

type Datadog struct {
	options *Options
}

func NewDatadogWithConfigPath(path string) (*Datadog, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func (d *Datadog) Register(ctx context.Context) (*chi.Config, error) {
	if !d.options.Enabled || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			c.Middleware(c.WithServiceName(datadog.Service())),
		},
	}, nil
}

func Register(ctx context.Context) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	d := NewDatadogWithOptions(o)
	return d.Register(ctx)
}