package opentelemetry

import (
	"net"
	"os"

	"github.com/americanas-go/config"
)

type Options struct {
	Enabled bool
	Service string
	Env     string
	Version string
	Host    string
	Port    string
	Addr    string
	Tags    map[string]string
}

// NewOptionsWithPath unmarshals options based on a given key path.
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

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {

	opts := &Options{}

	err := config.UnmarshalWithPath(root, opts)
	if err != nil {
		return nil, err
	}

	if v := os.Getenv("DD_SERVICE"); v != "" {
		opts.Service = v
	}

	if v := os.Getenv("DD_AGENT_HOST"); v != "" {
		opts.Host = v
	}

	if v := os.Getenv("DD_TRACE_AGENT_PORT"); v != "" {
		opts.Port = v
	}

	if v := os.Getenv("DD_ENV"); v != "" {
		opts.Env = v
	}

	if v := os.Getenv("DD_VERSION"); v != "" {
		opts.Version = v
	}

	opts.Addr = net.JoinHostPort(opts.Host, opts.Port)

	return opts, nil
}
