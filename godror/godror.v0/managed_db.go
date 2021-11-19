package godror

import (
	"context"
	"database/sql"
)

// ManagedDB represents a managed db client for oracle.
type ManagedDB struct {
	DB      *sql.DB
	Plugins []Plugin
	Options *Options
}

// NewManagedDBWithConfigPath returns a managed db client with options from config path.
func NewManagedDBWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*ManagedDB, error) {

	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewManagedDBWithOptions(ctx, opts, plugins...)
}

// NewManagedDB returns a managed db client with default options.
func NewManagedDB(ctx context.Context, plugins ...Plugin) (*ManagedDB, error) {
	opts, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewManagedDBWithOptions(ctx, opts, plugins...)
}

// NewManagedDBWithOptions returns a managed db client with options.
func NewManagedDBWithOptions(ctx context.Context, opts *Options, plugins ...Plugin) (*ManagedDB, error) {
	s, err := NewDBWithOptions(ctx, opts, plugins...)
	if err != nil {
		return nil, err
	}

	return &ManagedDB{
		DB:      s,
		Plugins: plugins,
		Options: opts,
	}, nil
}
