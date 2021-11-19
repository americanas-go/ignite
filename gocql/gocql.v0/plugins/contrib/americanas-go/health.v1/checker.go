package health

import (
	"context"

	"github.com/gocql/gocql"
)

// Checker represents a health checker for gocql cassandra client.
type Checker struct {
	session *gocql.Session
}

// Check checks if cassandra cluster is up.
func (c *Checker) Check(ctx context.Context) error {
	return c.session.Query("void").Exec()
}

// NewChecker returns a new health checker.
func NewChecker(session *gocql.Session) *Checker {
	return &Checker{session: session}
}
