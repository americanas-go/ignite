package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewHandler(schema *graphql.Schema) *handler.Handler {
	c, _ := DefaultHandlerConfig()
	return NewHandlerWithConfig(schema, c)
}

func NewHandlerWithConfig(schema *graphql.Schema, c *handler.Config) *handler.Handler {
	c.Schema = schema
	return handler.New(c)
}
