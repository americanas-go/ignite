package health

import (
	"context"

	"github.com/go-redis/redis/v9"
)

// ClientHealth represents redis cluster client checker.
type ClusterClientChecker struct {
	client *redis.ClusterClient
}

// Check  checks if redis cluster is responding.
func (c *ClusterClientChecker) Check(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

// NewClusterClientChecker creates the redis cluster client checker.
func NewClusterClientChecker(client *redis.ClusterClient) *ClusterClientChecker {
	return &ClusterClientChecker{client: client}
}
