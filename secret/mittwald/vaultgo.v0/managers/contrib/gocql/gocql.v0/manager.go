package gocql

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/db/nosql/gocql/gocql.v0"
	vault "github.com/americanas-go/ignite/secret/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

// Manager represents a vault manager for cassandra client.
type Manager struct {
	managedSession *gocql.ManagedSession
	options        *vault.ManagerOptions
}

// NewManager returns a new vault manager with default options.
func NewManager(managedSession *gocql.ManagedSession) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewManagerWithOptions(managedSession, o)
}

// NewManagerWithConfigPath returns a new vault manager with options from config path.
func NewManagerWithConfigPath(managedSession *gocql.ManagedSession, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewManagerWithOptions(managedSession, o), nil
}

// NewManagerWithOptions returns a new vault manager with options.
func NewManagerWithOptions(managedSession *gocql.ManagedSession, options *vault.ManagerOptions) vault.Manager {
	return &Manager{options: options, managedSession: managedSession}
}

// Options returns vault manager options.
func (m *Manager) Options() *vault.ManagerOptions {
	return m.options
}

// Close closes cassandra client.
func (m *Manager) Close(ctx context.Context) error {
	m.managedSession.Session.Close()
	return nil
}

// Configure configures cassandra client.
func (m *Manager) Configure(ctx context.Context, data map[string]interface{}) error {
	var username, password string
	var ok bool

	if username, ok = data["username"].(string); !ok {
		return errors.Internalf("username not found in data map")
	}

	if password, ok = data["password"].(string); !ok {
		return errors.Internalf("password not found in data map")
	}

	m.managedSession.Options.Username = username
	m.managedSession.Options.Password = password

	session, err := gocql.NewSessionWithOptions(ctx, m.managedSession.Options, m.managedSession.Plugins...)
	if err != nil {
		return err
	}

	m.managedSession.Session = session

	return nil
}
