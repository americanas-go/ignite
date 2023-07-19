package otelresty

import (
	"github.com/americanas-go/ignite"
	dubresty "github.com/dubonzi/otelresty"
)

type Options struct {
	Enabled        bool
	TracerName     string
	TracingOptions []dubresty.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions(tracingOptions ...dubresty.Option) (*Options, error) {
	opts := &Options{
		TracingOptions: tracingOptions,
	}

	return ignite.MergeOptionsWithPath[Options](opts, root)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, tracingOptions ...dubresty.Option) (opts *Options, err error) {
	opts, err = NewOptions(tracingOptions...)
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
