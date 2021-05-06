package health

import (
	"context"
	"database/sql"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
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

func (i *Health) Register(ctx context.Context, db *sql.DB) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating godror in health")

	checker := NewChecker(db)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("godror successfully integrated in health")

	return nil
}
