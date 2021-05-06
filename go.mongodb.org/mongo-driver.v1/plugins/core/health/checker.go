package health

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Checker struct {
	client *mongo.Client
}

func (c *Checker) Check(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

func NewChecker(client *mongo.Client) *Checker {
	return &Checker{client: client}
}
