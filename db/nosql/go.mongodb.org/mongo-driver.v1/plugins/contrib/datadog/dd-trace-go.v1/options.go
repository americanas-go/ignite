package datadog

import (
	"github.com/americanas-go/config"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

// Options represents datadog plugin for mongo options.
type Options struct {
	Enabled bool
	Options []mongotrace.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions(options ...mongotrace.Option) (*Options, error) {
	o := &Options{
		Options: options,
	}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string, options ...mongotrace.Option) (opts *Options, err error) {
	opts, err = NewOptions(options...)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
