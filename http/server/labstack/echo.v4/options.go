package echo

import "github.com/americanas-go/config"

// Options echo server options.
type Options struct {
	HideBanner   bool
	DisableHTTP2 bool `config:"disableHTTP2"`
	Port         int
	Type         string
	Protocol     string
	TLS          struct {
		Enabled bool
		Type    string
		Auto    struct {
			Host string
		}
		File struct {
			Cert string
			Key  string
		}
	} `config:"tls"`
	Json struct {
		Pretty struct {
			Enabled bool
		}
	}
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
