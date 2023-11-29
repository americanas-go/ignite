package otel

import (
	"os"

	"github.com/americanas-go/config"
)

type Options struct {
	Enabled  bool
	Service  string
	Env      string
	Version  string
	Protocol string
	Endpoint string
	Insecure bool
	Tags     map[string]string
	TLS      struct {
		Cert string
	}
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

	if v := os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL"); v != "" {
		opts.Protocol = v
	}

	if v := os.Getenv("OTEL_SERVICE_NAME"); v != "" {
		opts.Service = v
	}

	if v := os.Getenv("OTEL_SERVICE_VERSION"); v != "" {
		opts.Version = v
	}

	if v := os.Getenv("OTEL_ENV"); v != "" {
		opts.Env = v
	}

	return opts, nil
}
