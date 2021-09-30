package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type ManagedClient struct {
	Client  *redis.Client
	Plugins []Plugin
	Options *Options
}

func NewManagedClientWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*ManagedClient, error) {

	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	client, err := NewClientWithOptions(ctx, opts, plugins...)
	if err != nil {
		return nil, err
	}

	return &ManagedClient{
		Client:  client,
		Plugins: plugins,
		Options: opts,
	}, nil
}

func NewManagedClient(ctx context.Context, plugins ...Plugin) (*ManagedClient, error) {
	opts, err := NewOptions()
	if err != nil {
		return nil, err
	}

	client, err := NewClientWithOptions(ctx, opts, plugins...)
	if err != nil {
		return nil, err
	}

	return &ManagedClient{
		Client:  client,
		Plugins: plugins,
		Options: opts,
	}, nil
}

func NewManagedClientWithOptions(ctx context.Context, opts *Options, plugins ...Plugin) (*ManagedClient, error) {
	s, err := NewClientWithOptions(ctx, opts, plugins...)
	if err != nil {
		return nil, err
	}

	return &ManagedClient{
		Client:  s,
		Plugins: plugins,
		Options: opts,
	}, nil
}