package health

import (
	"github.com/americanas-go/config"
)

type Options struct {
	Name        string
	Host        string
	Endpoint    string
	Enabled     bool
	Description string
	Required    bool
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

// NewOptionsWithPath unmarshals options based a given key path.
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
