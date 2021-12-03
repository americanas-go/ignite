package godror

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/godror/godror.v0"
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

// Manager represents a vault manager for oracle client.
type Manager struct {
	managedDB *godror.ManagedDB
	options   *vault.ManagerOptions
}

// NewManager returns a new vault manager with default options.
func NewManager(managedDB *godror.ManagedDB) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewManagerWithOptions(managedDB, o)
}

// NewManagerWithConfigPath returns a new vault manager with options from config path.
func NewManagerWithConfigPath(managedDB *godror.ManagedDB, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewManagerWithOptions(managedDB, o), nil
}

// NewManagerWithOptions returns a new vault manager with options.
func NewManagerWithOptions(managedDB *godror.ManagedDB, options *vault.ManagerOptions) vault.Manager {
	return &Manager{options: options, managedDB: managedDB}
}

// Options returns vault manager options.
func (m *Manager) Options() *vault.ManagerOptions {
	return m.options
}

// Close closes oracle client.
func (m *Manager) Close(ctx context.Context) error {
	return m.managedDB.DB.Close()
}

// Configure configures oracle client.
func (m *Manager) Configure(ctx context.Context, data map[string]interface{}) error {
	var username, password string
	var ok bool

	if username, ok = data["username"].(string); !ok {
		return errors.Internalf("username not found in data map")
	}

	if password, ok = data["password"].(string); !ok {
		return errors.Internalf("password not found in data map")
	}

	m.managedDB.Options.Username = username
	m.managedDB.Options.Password = password

	db, err := godror.NewDBWithOptions(ctx, m.managedDB.Options, m.managedDB.Plugins...)
	if err != nil {
		return err
	}

	m.managedDB.DB = db

	return nil
}
