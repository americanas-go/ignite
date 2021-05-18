package health

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
)

type Checker struct {
	client *elasticsearch.Client
}

func (c *Checker) Check(ctx context.Context) error {
	_, err := c.client.Ping(c.client.Ping.WithPretty())
	return err
}

func NewChecker(client *elasticsearch.Client) *Checker {
	return &Checker{client: client}
}
