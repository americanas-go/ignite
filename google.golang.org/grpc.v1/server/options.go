package server

import (
	"github.com/americanas-go/ignite"
)

// Options grpc server options.
type Options struct {
	Port                  int
	MaxConcurrentStreams  int64
	InitialWindowSize     int32
	InitialConnWindowSize int32
	TLS                   TLSOptions `config:"tls"`
}

type TLSAutoOptions struct {
	Host string
}

type TLSOptions struct {
	Enabled bool
	Type    string
	Auto    TLSAutoOptions
	File    TLSFileOptions
}

type TLSFileOptions struct {
	Cert string
	Key  string
	CA   string `config:"ca"`
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}

return ignite.NewOptionsWithPath[Options](root, path)
}
