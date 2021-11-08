package redis

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/go-redis/redis.v7"
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

type ClusterManager struct {
	managedClusterClient *redis.ManagedClusterClient
	options              *vault.ManagerOptions
}

func NewClusterManager(managedClient *redis.ManagedClusterClient) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClusterManagerWithOptions(managedClient, o)
}

func NewClusterManagerWithConfigPath(managedClient *redis.ManagedClusterClient, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClusterManagerWithOptions(managedClient, o), nil
}

func NewClusterManagerWithOptions(managedClient *redis.ManagedClusterClient, options *vault.ManagerOptions) vault.Manager {
	return &ClusterManager{options: options, managedClusterClient: managedClient}
}

func (m *ClusterManager) Options() *vault.ManagerOptions {
	return m.options
}

func (m *ClusterManager) Close(ctx context.Context) error {
	return m.managedClusterClient.Client.Close()
}

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
