package resty

import (
	"time"

	"github.com/americanas-go/ignite"
)

// Options represents resty client options.
type Options struct {
	Debug             bool
	ConnectionTimeout time.Duration
	CloseConnection   bool
	KeepAlive         time.Duration
	RequestTimeout    time.Duration
	FallbackDelay     time.Duration
	Transport         OptionsTransport
	Host              string
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
