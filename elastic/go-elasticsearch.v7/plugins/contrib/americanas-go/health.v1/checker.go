package health

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
)

// Checker represents elasticsearch health check.
type Checker struct {
	client *elasticsearch.Client
}

// Check checks if elasticsearch is responding.
func (c *Checker) Check(ctx context.Context) error {
	_, err := c.client.Ping(c.client.Ping.WithPretty())
	return err
}

// NewChecker creates the elasticsearch health check.
func NewChecker(client *elasticsearch.Client) *Checker {
	return &Checker{client: client}
}
