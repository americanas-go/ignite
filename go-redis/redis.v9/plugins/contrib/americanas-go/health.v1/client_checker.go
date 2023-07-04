package health

import (
	"context"

	"github.com/go-redis/redis/v9"
)

// ClientChecker represents redis client checker.
type ClientChecker struct {
	client *redis.Client
}

// Check checks if redis server is responding.
func (c *ClientChecker) Check(ctx context.Context) error {
	return c.client.Conn().Ping(ctx).Err()
}

// NewClientChecker creates the redis client checker.
func NewClientChecker(client *redis.Client) *ClientChecker {
	return &ClientChecker{client: client}
}
