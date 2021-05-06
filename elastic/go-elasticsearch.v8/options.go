package elasticsearch

import (
	"time"

	"github.com/americanas-go/config"
)

type Options struct {
	Addresses             []string
	Username              string
	Password              string
	CloudID               string `config:"cloudID"`
	APIKey                string `config:"APIKey"`
	CACert                string `config:"CACert"`
	RetryOnStatus         []int
	DisableRetry          bool
	EnableRetryOnTimeout  bool
	MaxRetries            int
	DiscoverNodesOnStart  bool
	DiscoverNodesInterval time.Duration
	EnableMetrics         bool
	EnableDebugLogger     bool
	RetryBackoff          time.Duration
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
