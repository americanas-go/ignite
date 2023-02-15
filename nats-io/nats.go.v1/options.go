package nats

import (
	"time"

	"github.com/americanas-go/ignite"
)

// Options nats connection options.
type Options struct {
	Url           string
	MaxReconnects int
	ReconnectWait time.Duration
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}