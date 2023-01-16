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

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
