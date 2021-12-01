package datadog

import (
	"github.com/americanas-go/config"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

// Options datadog plugin for grpc server options.
type Options struct {
	Enabled      bool
	traceOptions []grpctrace.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions(traceOptions ...grpctrace.Option) (*Options, error) {
	o := &Options{
		traceOptions: traceOptions,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, traceOptions ...grpctrace.Option) (opts *Options, err error) {

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
