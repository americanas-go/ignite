package cors

import (
	"github.com/americanas-go/config"
)

// Options CORS plugin for echo server options.
type Options struct {
	Enabled bool
	Allowed struct {
		Origins     []string
		Headers     []string
		Methods     []string
		Credentials bool
	}
	Exposed struct {
		Headers []string
	}
	MaxAge int
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
