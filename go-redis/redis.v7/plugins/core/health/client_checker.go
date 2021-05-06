package health

import (
	"context"

	"github.com/go-redis/redis/v7"
)

type ClientChecker struct {
	client *redis.Client
}

func (c *ClientChecker) Check(ctx context.Context) error {
	return c.client.Conn().Ping().Err()
}

func NewClientChecker(client *redis.Client) *ClientChecker {
	return &ClientChecker{client: client}
}
