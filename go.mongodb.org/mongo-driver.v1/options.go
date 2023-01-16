package mongo

import (
	"github.com/americanas-go/ignite"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Options represents mongo client options.
type Options struct {
	Uri  string
	Auth *options.Credential
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
