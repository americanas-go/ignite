package client

import (
	"time"

	"github.com/americanas-go/config"
)

// Options grpc client options.
type Options struct {
	TLS                   TLSOptions `config:"tls"`
	InitialWindowSize     int32
	InitialConnWindowSize int32
	Host                  string
	Block                 bool
	HostOverwrite         string
	Port                  int
	Keepalive             KeepAliveOptions
	ConnectParams         struct {
		Backoff struct {
			BaseDelay  time.Duration
			Multiplier float64
			Jitter     float64
			MaxDelay   time.Duration
		}
		MinConnectTimeout time.Duration
	}
}

type TLSOptions struct {
	Enabled            bool
	CertFile           string
	KeyFile            string
	CAFile             string `config:"caFile"`
	InsecureSkipVerify bool
}

type KeepAliveOptions struct {
	Time                time.Duration
	Timeout             time.Duration
	PermitWithoutStream bool
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
