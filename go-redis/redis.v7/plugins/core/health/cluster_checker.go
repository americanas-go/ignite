package health

import (
	"context"

	"github.com/go-redis/redis/v7"
)

type ClusterClientChecker struct {
	client *redis.ClusterClient
}

func (c *ClusterClientChecker) Check(ctx context.Context) error {
	return c.client.Ping().Err()
}

func NewClusterClientChecker(client *redis.ClusterClient) *ClusterClientChecker {
	return &ClusterClientChecker{client: client}
}
