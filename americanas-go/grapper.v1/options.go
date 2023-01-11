package grapper

import (
	"github.com/americanas-go/config"
)

// Options struct which represents cors plugin from chi options.
type Options struct {
	Name bool
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath returns options from config path.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
