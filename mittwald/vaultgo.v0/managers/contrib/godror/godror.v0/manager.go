package godror

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/godror/godror.v0"
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

type Manager struct {
	managedDB *godror.ManagedDB
	options   *vault.ManagerOptions
}

func NewManager(managedDB *godror.ManagedDB) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewManagerWithOptions(managedDB, o)
}

func NewManagerWithConfigPath(managedDB *godror.ManagedDB, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewManagerWithOptions(managedDB, o), nil
}

func NewManagerWithOptions(managedDB *godror.ManagedDB, options *vault.ManagerOptions) vault.Manager {
	return &Manager{options: options, managedDB: managedDB}
}

func (m *Manager) Options() *vault.ManagerOptions {
	return m.options
}

func (m *Manager) Close(ctx context.Context) error {
	return m.managedDB.DB.Close()
}

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
