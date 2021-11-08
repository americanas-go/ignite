package mongo

import (
	"context"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	vault "github.com/americanas-go/ignite/mittwald/vaultgo.v0"
	"github.com/americanas-go/log"
)

type Manager struct {
	conn    *mongo.Conn
	options *vault.ManagerOptions
}

func NewManager(conn *mongo.Conn) vault.Manager {
	o, err := vault.NewManagerOptionsWithPath(root)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewManagerWithOptions(conn, o)
}

func NewManagerWithConfigPath(conn *mongo.Conn, path string) (vault.Manager, error) {
	o, err := vault.NewManagerOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewManagerWithOptions(conn, o), nil
}

func NewManagerWithOptions(conn *mongo.Conn, options *vault.ManagerOptions) vault.Manager {
	return &Manager{options: options, conn: conn}
}

func (m *Manager) Options() *vault.ManagerOptions {
	return m.options
}

func (m *Manager) Close(ctx context.Context) error {
	return m.conn.Client.Disconnect(ctx)
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

	m.conn.Options.Auth.Username = username
	m.conn.Options.Auth.Password = password

	conn, err := mongo.NewConnWithOptions(ctx, m.conn.Options, m.conn.Plugins...)
	if err != nil {
		return err
	}

	m.conn.Database = conn.Database
	m.conn.Client = conn.Client

	return nil
}
