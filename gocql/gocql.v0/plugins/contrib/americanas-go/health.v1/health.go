package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/gocql/gocql"
)

func Register(ctx context.Context, session *gocql.Session) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewHealthWithOptions(o)
	return h.Register(ctx, session)
}

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

func (i *Health) Register(ctx context.Context, session *gocql.Session) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating gocql in health")

	checker := NewChecker(session)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("gocql successfully integrated in health")

	return nil
}
