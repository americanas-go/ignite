package aws

import (
	"os"
	"time"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/net/http/client"
)

type OptionsCustomEndpoint map[string]struct {
	PartitionID       string `config:"partitionId"`
	URL               string `config:"url"`
	SigningRegion     string
	HostnameImmutable bool
}

type Options struct {
	AccessKeyId                 string
	SecretAccessKey             string
	DefaultRegion               string
	SessionToken                string
	DefaultAccountNumber        string
	MaxAttempts                 int
	CustomEndpoint              OptionsCustomEndpoint
	HasRateLimit                bool
	MaxConnsPerHost             int
	MaxIdleConns                int
	MaxIdleConnsPerHost         int
	TimeoutMillis               time.Duration
	KeepAliveMillis             time.Duration
	IdleConnTimeoutMillis       time.Duration
	ResponseHeaderTimeoutMillis time.Duration
	HttpClient                  client.Options
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	opts, err := ignite.NewOptionsWithPath[Options](root)
	if err != nil {
		return nil, err
	}

	if v := os.Getenv("AWS_ACCESS_KEY_ID"); v != "" {
		opts.AccessKeyId = v
	}

	if v := os.Getenv("AWS_SECRET_ACCESS_KEY"); v != "" {
		opts.SecretAccessKey = v
	}

	if v := os.Getenv("AWS_DEFAULT_REGION"); v != "" {
		opts.DefaultRegion = v
	}

	if v := os.Getenv("AWS_DEFAULT_ACCOUNT_NUMBER"); v != "" {
		opts.DefaultAccountNumber = v
	}

	if v := os.Getenv("AWS_SESSION_TOKEN"); v != "" {
		opts.SessionToken = v
	}

	return opts, nil
}
