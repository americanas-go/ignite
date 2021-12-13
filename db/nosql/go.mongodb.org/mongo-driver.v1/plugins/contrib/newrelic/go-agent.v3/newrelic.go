package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Newrelic represents newrelic plugin for mongo.
type Newrelic struct {
	options *Options
}

// NewNewrelicWithConfigPath returns a new newrelic plugin with options from config path.
func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

// NewNewrelicWithOptions returns a new newrelic plugin with options.
func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

// Register registers this newrelic plugin on a new mongo client.
func (d *Newrelic) Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {

	if !d.options.Enabled || !newrelic.IsEnabled() {
		return nil, nil
	}

	return func(ctx context.Context, options *options.ClientOptions) error {
		logger := log.FromContext(ctx)

		logger.Trace("integrating mongo in newrelic")

		options.SetMonitor(nrmongo.NewCommandMonitor(options.Monitor))

		logger.Debug("mongo successfully integrated in newrelic")

		return nil
	}, nil
}

// Register registers a new newrelic plugin on a new mongo client.
func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	datadog := NewNewrelicWithOptions(o)
	return datadog.Register(ctx)
}
