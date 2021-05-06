package health

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
)

type Checker struct {
	conn *nats.Conn
}

func (c *Checker) Check(ctx context.Context) error {

	var err error

	if !c.conn.IsConnected() {
		err = errors.New("Not connected")
	}

	return err
}

func NewChecker(conn *nats.Conn) *Checker {
	return &Checker{conn: conn}
}
