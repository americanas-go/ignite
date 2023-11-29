package contrib

import (
	"context"

	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/ignite/go.opentelemetry.io/otel.v1"
	"github.com/americanas-go/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// OtelMongo represents a opentelemetry plugin for mongo.
type OtelMongo struct {
	options *Options
}

// NewOtelMongoWithConfigPath returns a new opentelemetry plugin with options from config path.
func NewOtelMongoWithConfigPath(path string) (*OtelMongo, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOtelMongoWithOptions(o), nil
}

// NewOtelMongoWithOptions returns a new opentelemetry plugin with options.
func NewOtelMongoWithOptions(options *Options) *OtelMongo {
	return &OtelMongo{options: options}
}

// NewOtelMongo returns a new opentelemetry plugin with default options.
func NewOtelMongo() *OtelMongo {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOtelMongoWithOptions(o)
}

// Register registers this opentelemetry plugin on a new mongo client.
func (d *OtelMongo) Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	if !d.options.Enabled || !otel.IsTracerEnabled() {
		return nil, nil
	}

	return func(ctx context.Context, options *options.ClientOptions) error {
		logger := log.FromContext(ctx)

		logger.Trace("integrating opentelemetry in mongo")

		options.SetMonitor(otelmongo.NewMonitor())

		logger.Debug("opentelemetry successfully integrated in mongo")

		return nil
	}, nil
}

// Register registers a new opentelemetry plugin on a new mongo client.
func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	plugin := NewOtelMongoWithOptions(o)
	return plugin.Register(ctx)
}
