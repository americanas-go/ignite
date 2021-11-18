package elasticsearch

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
)

// ManagedClient struct that represents a managed client.
type ManagedClient struct {
	Client  *elasticsearch.Client
	Plugins []Plugin
	Options *Options
}

// NewManagedClientWithConfigPath returns a managed client with options from config path.
func NewManagedClientWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*ManagedClient, error) {

	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewManagedClientWithOptions(ctx, opts, plugins...)
}

// NewManagedClientWithConfigPath returns a managed client with default options.
func NewManagedClient(ctx context.Context, plugins ...Plugin) (*ManagedClient, error) {
	opts, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewManagedClientWithOptions(ctx, opts, plugins...)
}

// NewManagedClientWithConfigPath returns a managed client with options.
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
