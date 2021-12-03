package health

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Checker health checker for grpc client.
type Checker struct {
	conn *grpc.ClientConn
}

// Check check if grpc client is connected.
func (c *Checker) Check(ctx context.Context) error {

	var err error

	if c.conn.GetState() != connectivity.Ready {
		err = errors.New("not ready")
	}

	return err
}

// NewChecker returns a new health checker for grpc client.
func NewChecker(conn *grpc.ClientConn) *Checker {
	return &Checker{conn: conn}
}
