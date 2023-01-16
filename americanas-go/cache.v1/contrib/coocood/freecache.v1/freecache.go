package freecache

import (
	"github.com/americanas-go/cache"
	cfreecache "github.com/americanas-go/cache/driver/contrib/coocood/freecache.v1"
	"github.com/coocood/freecache"
)

// NewDriverWithConfigPath returns a cache with options from config path .
func NewDriverWithConfigPath(cache *freecache.Cache, path string) (cache.Driver, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return cfreecache.New(cache, options), nil
}

// NewDriver returns a cache.
func NewDriver(cache *freecache.Cache) (c cache.Driver, err error) {
	var options *cfreecache.Options
	options, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return cfreecache.New(cache, options), nil
}
