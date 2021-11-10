package datadog

import (
	"github.com/americanas-go/config"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

type Options struct {
	Enabled       bool
	OperationName string
	SpanOptions   []ddtrace.StartSpanOption
}

// NewOptions returns options from config file or environment vars.
func NewOptions(spanOptions ...ddtrace.StartSpanOption) (*Options, error) {
	o := &Options{
		SpanOptions: spanOptions,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, spanOptions ...ddtrace.StartSpanOption) (opts *Options, err error) {
	opts, err = NewOptions(spanOptions...)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
