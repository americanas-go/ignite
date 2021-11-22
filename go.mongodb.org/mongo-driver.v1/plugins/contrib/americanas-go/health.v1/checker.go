package health

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Checker represents health checker plugin for mongoDB.
type Checker struct {
	client *mongo.Client
}

// Check checks if mongo is up (no returned error).
func (c *Checker) Check(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

// NewChecker returns a new checker plugin.
func NewChecker(client *mongo.Client) *Checker {
	return &Checker{client: client}
}
