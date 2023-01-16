package echo

import (
	"github.com/americanas-go/ignite"
)

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
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}