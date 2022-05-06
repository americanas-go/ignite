package gocql

import (
	"time"

	"github.com/americanas-go/config"
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
