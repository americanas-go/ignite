package redis

import (
	"context"

	"github.com/go-redis/redis/v7"
)

type ManagedClusterClient struct {
	Client  *redis.ClusterClient
	Plugins []ClusterPlugin
	Options *Options
}

func NewManagedClusterClientWithConfigPath(ctx context.Context, path string, plugins ...ClusterPlugin) (*ManagedClusterClient, error) {

	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewManagedClusterClientWithOptions(ctx, opts, plugins...)
}

func NewManagedClusterClient(ctx context.Context, plugins ...ClusterPlugin) (*ManagedClusterClient, error) {
	opts, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewManagedClusterClientWithOptions(ctx, opts, plugins...)
}

func NewManagedClusterClientWithOptions(ctx context.Context, opts *Options, plugins ...ClusterPlugin) (*ManagedClusterClient, error) {
	s, err := NewClusterClientWithOptions(ctx, opts, plugins...)
	if err != nil {
		return nil, err
	}

	return &ManagedClusterClient{
		Client:  s,
		Plugins: plugins,
		Options: opts,
	}, nil
}
