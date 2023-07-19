package otelresty

import (
	"context"

	"github.com/americanas-go/ignite/opentelemetry/opentelemetry-go.v1"
	"github.com/americanas-go/log"
	dubresty "github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
)

// Otelresty represents the Opentelemetry integration for resty.
type Otelresty struct {
	options *Options
}

// NewOtelrestyWithConfigPath returns a new Otelresty with options from the provided path.
func NewOtelrestyWithConfigPath(path string, tracingOptions ...dubresty.Option) (*Otelresty, error) {
	o, err := NewOptionsWithPath(path, tracingOptions...)
	if err != nil {
		return nil, err
	}
	return NewOtelrestyWithOptions(o), nil
}

// NewOtelresty returns a new Otelresty with default options.
func NewOtelresty(tracingOptions ...dubresty.Option) *Otelresty {
	o, err := NewOptions(tracingOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOtelrestyWithOptions(o)
}

// NewOtelresty returns a new Otelresty with default options.
func NewOtelrestyWithOptions(options *Options) *Otelresty {
	return &Otelresty{options: options}
}

// Register registers the Opentelemetry integration with the provided resty client. It is a shorthand for NewOtelresty().Register().
func Register(ctx context.Context, client *resty.Client) error {
	options, err := NewOptions()
	if err != nil {
		return err
	}
	o := NewOtelrestyWithOptions(options)
	return o.Register(ctx, client)
}

func (o *Otelresty) Register(ctx context.Context, client *resty.Client) error {
	if !o.options.Enabled || !opentelemetry.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty with opentelemetry")

	o.options.TracingOptions = append(o.options.TracingOptions, dubresty.WithTracerName(o.options.TracerName))

	dubresty.TraceClient(client, o.options.TracingOptions...)

	logger.Debug("resty successfully integrated with opentelemetry")
	return nil
}
