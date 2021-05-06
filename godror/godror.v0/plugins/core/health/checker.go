package health

import (
	"context"
	"database/sql"
)

type Checker struct {
	db *sql.DB
}

func (c *Checker) Check(ctx context.Context) error {
	if err := c.db.Ping(); err != nil {
		return err
	}
	return nil
}

func NewChecker(db *sql.DB) *Checker {
	return &Checker{db: db}
}
