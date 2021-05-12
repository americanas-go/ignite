package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/go.mongodb.org/mongo-driver.v1"
	m "go.mongodb.org/mongo-driver/mongo"
)

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

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
