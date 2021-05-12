package datadog

import (
	"net"
	"os"

	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/net/http/client"
)

type Options struct {
	Service       string
	Env           string
	Enabled       bool
	Tags          map[string]string
	Host          string
	Port          string
	LambdaMode    bool
	Analytics     bool
	AnalyticsRate float64
	DebugMode     bool
	DebugStack    bool
	HttpClient    client.Options
	Version       string
	Log           struct {
		Level string
	}
	Addr string
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
