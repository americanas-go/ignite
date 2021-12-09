package datadog

import (
	"github.com/americanas-go/config"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

// Options represents datadog plugin for go driver for oracle options.
type Options struct {
	Enabled      bool
	TraceOptions []sqltrace.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions(traceOptions ...sqltrace.Option) (*Options, error) {
	o := &Options{
		TraceOptions: traceOptions,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, traceOptions ...sqltrace.Option) (opts *Options, err error) {
	opts, err = NewOptions(traceOptions...)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
