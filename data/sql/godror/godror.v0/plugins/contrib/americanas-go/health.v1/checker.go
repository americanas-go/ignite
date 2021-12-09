package health

import (
	"context"
	"database/sql"
)

// Checker health checker for go driver for oracle.
type Checker struct {
	db *sql.DB
}

// Check checks if db server is up.
func (c *Checker) Check(ctx context.Context) error {
	if err := c.db.Ping(); err != nil {
		return err
	}
	return nil
}

//NewChecker returns a new health checker.
func NewChecker(db *sql.DB) *Checker {
	return &Checker{db: db}
}
