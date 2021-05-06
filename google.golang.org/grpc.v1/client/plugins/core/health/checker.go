package health

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type Checker struct {
	conn *grpc.ClientConn
}

func (c *Checker) Check(ctx context.Context) error {

	var err error

	if c.conn.GetState() != connectivity.Ready {
		err = errors.New("not ready")
	}

	return err
}

func NewChecker(conn *grpc.ClientConn) *Checker {
	return &Checker{conn: conn}
}
