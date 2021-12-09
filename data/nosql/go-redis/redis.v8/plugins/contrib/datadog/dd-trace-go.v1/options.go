package datadog

import (
	"github.com/americanas-go/config"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
)

// Options represents a datadog client for redis options.
type Options struct {
	Enabled      bool
	TraceOptions []redistrace.ClientOption
}

// NewOptions returns options from config or environment vars.
func NewOptions(traceOptions ...redistrace.ClientOption) (*Options, error) {
	o := &Options{TraceOptions: traceOptions}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, traceOptions ...redistrace.ClientOption) (opts *Options, err error) {
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
