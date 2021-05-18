package health

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type ClientChecker struct {
	client *redis.Client
}

func (c *ClientChecker) Check(ctx context.Context) error {
	return c.client.Conn(ctx).Ping(ctx).Err()
}

func NewClientChecker(client *redis.Client) *ClientChecker {
	return &ClientChecker{client: client}
}
