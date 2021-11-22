package gocql

import (
	"context"

	"github.com/gocql/gocql"
)

// ManagedSession represents a gocqsl managed session
type ManagedSession struct {
	Session *gocql.Session
	Plugins []Plugin
	Options *Options
}

// NewManagedSessionWithConfigPath returns a ManagedSession with options from config path.
func NewManagedSessionWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*ManagedSession, error) {

	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewManagedSessionWithOptions(ctx, opts, plugins...)
}

// NewManagedSession returns a ManagedSession with default options.
func NewManagedSession(ctx context.Context, plugins ...Plugin) (*ManagedSession, error) {
	opts, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewManagedSessionWithOptions(ctx, opts, plugins...)
}

// NewManagedSessionWithOptions returns a ManagedSession with options.
func NewManagedSessionWithOptions(ctx context.Context, opts *Options, plugins ...Plugin) (*ManagedSession, error) {
	s, err := NewSessionWithOptions(ctx, opts, plugins...)
	if err != nil {
		return nil, err
	}

	return &ManagedSession{
		Session: s,
		Plugins: plugins,
		Options: opts,
	}, nil
}
