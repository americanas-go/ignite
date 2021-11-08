package mongo

import (
	"context"
	"strings"

	"github.com/americanas-go/log"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Conn struct {
	ClientOptions *options.ClientOptions
	Client        *mongo.Client
	Database      *mongo.Database
	Options       *Options
	Plugins       []Plugin
}

type ClientOptionsPlugin func(context.Context, *options.ClientOptions) error
type ClientPlugin func(context.Context, *mongo.Client) error

type Plugin func(context.Context) (ClientOptionsPlugin, ClientPlugin)

func NewConn(ctx context.Context, plugins ...Plugin) (*Conn, error) {

	logger := log.FromContext(ctx)

	o, err := NewOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewConnWithOptions(ctx, o, plugins...)
}

func NewConnWithConfigPath(ctx context.Context, path string, plugins ...Plugin) (*Conn, error) {
	opts, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewConnWithOptions(ctx, opts, plugins...)
}

func NewConnWithOptions(ctx context.Context, o *Options, plugins ...Plugin) (conn *Conn, err error) {

	logger := log.FromContext(ctx)

	var clientOptionsPlugins []ClientOptionsPlugin
	var clientPlugins []ClientPlugin

	for _, plugin := range plugins {
		clientOptionsPlugin, clientPlugin := plugin(ctx)
		if clientOptionsPlugin != nil {
			clientOptionsPlugins = append(clientOptionsPlugins, clientOptionsPlugin)
		}
		if clientPlugin != nil {
			clientPlugins = append(clientPlugins, clientPlugin)
		}
	}

	co := clientOptions(ctx, o)

	for _, clientOptionsPlugin := range clientOptionsPlugins {
		if err := clientOptionsPlugin(ctx, co); err != nil {
			logger.Fatalf(err.Error())
		}
	}

	var client *mongo.Client
	var database *mongo.Database

	client, database, err = newClient(ctx, co)
	if err != nil {
		return nil, err
	}

	for _, clientPlugin := range clientPlugins {
		if err := clientPlugin(ctx, client); err != nil {
			logger.Fatalf(err.Error())
		}
	}

	conn = &Conn{
		ClientOptions: co,
		Client:        client,
		Database:      database,
		Plugins:       plugins,
		Options:       o,
	}

	return conn, err
}

func newClient(ctx context.Context, co *options.ClientOptions) (client *mongo.Client, database *mongo.Database, err error) {

	logger := log.FromContext(ctx)

	client, err = mongo.Connect(ctx, co)

	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	var connFields connstring.ConnString

	connFields, err = connstring.Parse(co.GetURI())

	if err != nil {
		return nil, nil, err
	}

	database = client.Database(connFields.Database)

	logger.Infof("Connected to MongoDB server: %v", strings.Join(connFields.Hosts, ","))

	return client, database, err
}

func clientOptions(ctx context.Context, o *Options) *options.ClientOptions {

	logger := log.FromContext(ctx)

	clientOptions := options.Client().ApplyURI(o.Uri)
	clientOptions.SetMonitor(&event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			logger.Debugf("mongodb cmd - %v %s %s %v", startedEvent.ConnectionID, startedEvent.CommandName, startedEvent.DatabaseName, startedEvent.RequestID)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			logger.Debugf("mongodb cmd - %v %s %vus %v", succeededEvent.ConnectionID, succeededEvent.CommandName, succeededEvent.DurationNanos, succeededEvent.RequestID)
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			logger.Errorf("mongodb cmd - %v %s %s %v", failedEvent.ConnectionID, failedEvent.CommandName, failedEvent.Failure, failedEvent.RequestID)
		},
	})
	clientOptions.SetPoolMonitor(&event.PoolMonitor{
		Event: func(poolEvent *event.PoolEvent) {
			logger.Debugf("mongodb conn pool - %v %s %s %s", poolEvent.ConnectionID, poolEvent.Type, poolEvent.Reason, poolEvent.Address)
		},
	})

	if o.Auth != nil {
		setAuthOptions(o, clientOptions)
	}

	return clientOptions
}

func setAuthOptions(o *Options, clientOptions *options.ClientOptions) {

	if o.Auth.Password == "" && o.Auth.Username == "" {
		return
	}

	if clientOptions.Auth == nil {
		clientOptions.Auth = &options.Credential{}
	}

	if o.Auth.Password != "" {
		clientOptions.Auth.Password = o.Auth.Password
		clientOptions.Auth.PasswordSet = true
	}

	if o.Auth.Username != "" {
		clientOptions.Auth.Username = o.Auth.Username
	}

	if clientOptions.Auth.AuthSource == "" {
		connFields, _ := connstring.Parse(clientOptions.GetURI())
		clientOptions.Auth.AuthSource = connFields.Database
	}
}
