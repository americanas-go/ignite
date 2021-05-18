package health

import (
	"context"

	"github.com/gocql/gocql"
)

type Checker struct {
	session *gocql.Session
}

func (c *Checker) Check(ctx context.Context) error {
	return c.session.Query("void").Exec()
}

func NewChecker(session *gocql.Session) *Checker {
	return &Checker{session: session}
}
