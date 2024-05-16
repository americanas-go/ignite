package azblob

import (
	"github.com/americanas-go/ignite"
)

type Options struct {
	AccountName string
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
	return opts, nil
}
