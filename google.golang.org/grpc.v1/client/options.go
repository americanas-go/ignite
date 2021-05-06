package client

import (
	"time"

	"github.com/americanas-go/config"
)

type Options struct {
	TLS struct {
		Enabled            bool
		CertFile           string
		KeyFile            string
		CAFile             string `config:"caFile"`
		InsecureSkipVerify bool
	} `config:"tls"`
	InitialWindowSize     int32
	InitialConnWindowSize int32
	Host                  string
	Block                 bool
	HostOverwrite         string
	Port                  int
	Keepalive             struct {
		Time                time.Duration
		Timeout             time.Duration
		PermitWithoutStream bool
	}
	ConnectParams struct {
		Backoff struct {
			BaseDelay  time.Duration
			Multiplier float64
			Jitter     float64
			MaxDelay   time.Duration
		}
		MinConnectTimeout time.Duration
	}
}

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

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
