package datadog

import (
	"github.com/americanas-go/ignite"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

// Options represents datadog client for resty options.
type Options struct {
	Enabled       bool
	OperationName string
	SpanOptions   []ddtrace.StartSpanOption
}

// NewOptions returns options from config file or environment vars.
func NewOptions(spanOptions ...ddtrace.StartSpanOption) (*Options, error) {
	opts := &Options{
		SpanOptions: spanOptions,
	}

	return ignite.MergeOptionsWithPath[Options](opts, root)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, spanOptions ...ddtrace.StartSpanOption) (opts *Options, err error) {
	opts, err = NewOptions(spanOptions...)
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
