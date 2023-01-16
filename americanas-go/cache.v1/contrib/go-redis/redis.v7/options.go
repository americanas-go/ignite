package redis

import (
	"github.com/americanas-go/cache/driver/contrib/go-redis/redis.v7"
	"github.com/americanas-go/ignite"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*redis.Options, error) {
	return ignite.NewOptionsWithPath[redis.Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *redis.Options, err error) {
	return ignite.NewOptionsWithPath[redis.Options](root, path)
}
