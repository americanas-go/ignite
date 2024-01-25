package health

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
)

// Register registers a new health check on sql DB.
func Register(ctx context.Context, db *sql.DB, connector driver.Connector) (d *sql.DB, err error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	h := NewHealthWithOptions(o)
	return h.Register(ctx, db, connector)
}

// Health represents a health check for go driver for oracle.
type Health struct {
	options *Options
}

// NewHealthWithOptions returns a new health check with options.
func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

// NewHealthWithConfigPath returns a new health check with options from config path.
func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

// NewHealth returns a new health check.
func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

// Register registers this health check on sql DB.
func (i *Health) Register(ctx context.Context, db *sql.DB, connector driver.Connector) (d *sql.DB, err error) {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating godror in health")

	checker := NewChecker(db)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("godror successfully integrated in health")

	return db, nil
}
