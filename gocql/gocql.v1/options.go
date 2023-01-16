package gocql

import (
	"time"

	"github.com/americanas-go/ignite"
)

// Options represents gocql options.
type Options struct {
	Hosts                    []string
	Port                     int
	DC                       string `config:"dc"`
	Username                 string
	Password                 string
	CQLVersion               string `config:"CQLVersion"`
	ProtoVersion             int
	Timeout                  time.Duration
	ConnectTimeout           time.Duration
	Keyspace                 string
	NumConns                 int
	Consistency              string
	SocketKeepalive          time.Duration
	MaxPreparedStmts         int
	MaxRoutingKeyInfo        int
	PageSize                 int
	DefaultTimestamp         bool
	ReconnectInterval        time.Duration
	MaxWaitSchemaAgreement   time.Duration
	DisableInitialHostLookup bool
	WriteCoalesceWaitTime    time.Duration
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
