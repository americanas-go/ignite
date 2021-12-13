package redis

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/db/nosql/go-redis/redis.v8"
	vault "github.com/americanas-go/ignite/secret/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

// Manager represents a vault manager for redis cluster client.
type ClusterManager struct {
	managedClusterClient *redis.ManagedClusterClient
	options              *vault.ManagerOptions
}

// NewManager returns a new vault manager with default options.
func NewClusterManager(managedClient *redis.ManagedClusterClient) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClusterManagerWithOptions(managedClient, o)
}

// NewManagerWithConfigPath returns a new vault manager with options from config path.
func NewClusterManagerWithConfigPath(managedClient *redis.ManagedClusterClient, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClusterManagerWithOptions(managedClient, o), nil
}

// NewManagerWithOptions returns a new vault manager with options.
func NewClusterManagerWithOptions(managedClient *redis.ManagedClusterClient, options *vault.ManagerOptions) vault.Manager {
	return &ClusterManager{options: options, managedClusterClient: managedClient}
}

// Options returns vault manager options.
func (m *ClusterManager) Options() *vault.ManagerOptions {
	return m.options
}

// Close closes redis cluster client.
func (m *ClusterManager) Close(ctx context.Context) error {
	return m.managedClusterClient.Client.Close()
}

// Configure configures redis cluster client.
func (m *ClusterManager) Configure(ctx context.Context, data map[string]interface{}) error {
	var password string
	var ok bool

	if password, ok = data["password"].(string); !ok {
		return errors.Internalf("password not found in data map")
	}

	m.managedClusterClient.Options.Password = password

	client, err := redis.NewClusterClientWithOptions(ctx, m.managedClusterClient.Options, m.managedClusterClient.Plugins...)
	if err != nil {
		return err
	}

	m.managedClusterClient.Client = client

	return nil
}
