package freecache

import (
	"github.com/americanas-go/cache/driver/contrib/coocood/freecache.v1"
	"github.com/americanas-go/ignite"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*freecache.Options, error) {
	return ignite.NewOptionsWithPath[freecache.Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *freecache.Options, err error) {
	return ignite.NewOptionsWithPath[freecache.Options](root, path)
}
