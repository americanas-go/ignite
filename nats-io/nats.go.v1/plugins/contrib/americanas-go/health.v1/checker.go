package health

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
)

// Checker health checker for nats.
type Checker struct {
	conn *nats.Conn
}

// Check checks if nats connection is up.
func (c *Checker) Check(ctx context.Context) error {

	var err error

	if !c.conn.IsConnected() {
		err = errors.New("Not connected")
	}

	return err
}

// NewChecker returns a new health checker for nats.
func NewChecker(conn *nats.Conn) *Checker {
	return &Checker{conn: conn}
}
