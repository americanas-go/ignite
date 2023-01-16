package kafka

import "github.com/americanas-go/ignite"

// Options kafka connection options.
type Options struct {
	Address   string
	Topic     string
	Partition int
	Network   string
	ConnType  string
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}