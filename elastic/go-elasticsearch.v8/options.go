package elasticsearch

import (
	"time"

	"github.com/americanas-go/ignite"
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

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}
