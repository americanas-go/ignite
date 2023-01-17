package datadog

import (
	"net"
	"os"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/net/http/client"
)

type Options struct {
	Service string
	Env     string
	Tracer  struct {
		Enabled bool
	}
	Profiler struct {
		Enabled bool
	}
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

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	opts, err := ignite.NewOptionsWithPath[Options](root)
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
