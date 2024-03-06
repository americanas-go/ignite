package bigquery

import (
	"github.com/americanas-go/ignite"
)

// Options bigquery client options.
type Options struct {
	ProjectID   string `config:"projectId"`
	Credentials struct {
		File string
		JSON string `config:"json"`
	}
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}
