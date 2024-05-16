package azure

import (
	"github.com/americanas-go/ignite"
)

type Options struct {
	ClientId     string
	ClientSecret string
	TenantId     string
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](Root, path)
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	opts, err := ignite.NewOptionsWithPath[Options](Root)
	if err != nil {
		return nil, err
	}
	return opts, nil
}