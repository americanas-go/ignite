package redis

import (
	"github.com/americanas-go/cache"
	credis "github.com/americanas-go/cache/driver/contrib/go-redis/redis.v7"
	"github.com/go-redis/redis/v7"
)

// NewClusterDriverWithConfigPath returns a cache with options from config path .
func NewClusterDriverWithConfigPath(client *redis.ClusterClient, path string) (cache.Driver, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return credis.NewCluster(client, options), nil
}

// NewClusterDriver returns a cache.
func NewClusterDriver(client *redis.ClusterClient) (c cache.Driver, err error) {
	var options *credis.Options
	options, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return credis.NewCluster(client, options), nil
}
