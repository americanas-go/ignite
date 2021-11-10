package elasticsearch

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/elastic/go-elasticsearch.v8"
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

type Manager struct {
	managedClient *elasticsearch.ManagedClient
	options       *vault.ManagerOptions
}

func NewManager(managedClient *elasticsearch.ManagedClient) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewManagerWithOptions(managedClient, o)
}

func NewManagerWithConfigPath(managedClient *elasticsearch.ManagedClient, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewManagerWithOptions(managedClient, o), nil
}

func NewManagerWithOptions(managedClient *elasticsearch.ManagedClient, options *vault.ManagerOptions) vault.Manager {
	return &Manager{options: options, managedClient: managedClient}
}

func (m *Manager) Options() *vault.ManagerOptions {
	return m.options
}

func (m *Manager) Close(ctx context.Context) error {
	return nil
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

	m.managedClient.Options.Password = password
	m.managedClient.Options.Username = username

	client, err := elasticsearch.NewClientWithOptions(ctx, m.managedClient.Options, m.managedClient.Plugins...)
	if err != nil {
		return err
	}

	m.managedClient.Client = client

	return nil
}
