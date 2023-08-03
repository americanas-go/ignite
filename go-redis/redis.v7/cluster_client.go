package redis

import (
	"context"
	"strings"

	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
)

// ClusterPlugin represents a redis cluster plugin func signature.
type ClusterPlugin func(context.Context, *redis.ClusterClient) error

// NewClusterClient returns a new ClusterClient.
func NewClusterClient(ctx context.Context, plugins ...ClusterPlugin) (*redis.ClusterClient, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Errorf(err.Error())
	}

	return NewClusterClientWithOptions(ctx, o, plugins...)
}

// NewClusterClientWithConfigPath returns a new ClusterClient with options from config path.
func NewClusterClientWithConfigPath(ctx context.Context, path string, plugins ...ClusterPlugin) (*redis.ClusterClient, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClusterClientWithOptions(ctx, opts, plugins...)
}

// NewClusterClientWithOptions returns a new ClusterClient with options.
func NewClusterClientWithOptions(ctx context.Context, o *Options, plugins ...ClusterPlugin) (client *redis.ClusterClient, err error) {

	logger := log.FromContext(ctx)

	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              o.Cluster.Addrs,
		MaxRedirects:       o.Cluster.MaxRedirects,
		ReadOnly:           o.Cluster.ReadOnly,
		RouteByLatency:     o.Cluster.RouteByLatency,
		RouteRandomly:      o.Cluster.RouteRandomly,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})

	ping := client.Ping()
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	for _, plugin := range plugins {
		if err := plugin(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to Redis Cluster server: %s status: %s", strings.Join(client.Options().Addrs, ","), ping.String())

	return client, err
}
