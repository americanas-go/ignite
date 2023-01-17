package redis

import (
	"github.com/americanas-go/cache"
	credis "github.com/americanas-go/cache/driver/contrib/go-redis/redis.v7"
	"github.com/go-redis/redis/v7"
)

// NewClientDriverWithConfigPath returns a cache with options from config path .
func NewClientDriverWithConfigPath(client *redis.Client, path string) (cache.Driver, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return credis.NewClient(client, options), nil
}

// NewClientDriver returns a cache.
func NewClientDriver(client *redis.Client) (c cache.Driver, err error) {
	var options *credis.Options
	options, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return credis.NewClient(client, options), nil
}
