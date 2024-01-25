package datadog

import (
	"context"
	c "github.com/go-chi/chi/v5"
	"net/http"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	chitrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi.v5"
)

// Datadog struct that represents a datadog plugin for chi
type Datadog struct {
	options *Options
}

// NewDatadogWithConfigPath returns a new datadog plugin with options from config path.
func NewDatadogWithConfigPath(path string, traceOptions ...chitrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

// NewDatadogWithOptions returns a new datadog plugin with options.
func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

// NewDatadog returns a new datadog plugin with default options.
func NewDatadog(traceOptions ...chitrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

// Register registers the datadog plugin to a new chi config.
func (d *Datadog) Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	if !d.options.Enabled || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			chitrace.Middleware(d.options.TraceOptions...),
		},
	}, nil
}

// Register registers a default datadog plugin to a new chi config.
func Register(ctx context.Context, mux *c.Mux) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	d := NewDatadogWithOptions(o)
	return d.Register(ctx, mux)
}
