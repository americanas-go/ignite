package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/log"
	m "go.mongodb.org/mongo-driver/mongo"
)

// Health represents a healh plugin for mongo.
type Health struct {
	options *Options
}

// NewHealthWithOptions returns a new health with options.
func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

// NewHealthWithConfigPath returns a new health with options from config path.
func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

// NewHealth returns a new health with default options.
func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return NewHealthWithOptions(o)
}

// Register registers this health plugin on a new mongo client.
func (i *Health) Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {

	logger := log.WithTypeOf(*i)

	return nil, func(ctx context.Context, client *m.Client) error {

		logger.Trace("integrating mongo in health")

		checker := NewChecker(client)
		hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
		health.Add(hc)

		logger.Debug("mongo successfully integrated in health")

		return nil
	}
}

// Register registers a new health plugin on a new mongo client.
func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewHealthWithOptions(o)
	return h.Register(ctx)
}
