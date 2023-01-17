package graphql

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite"
	"github.com/graphql-go/handler"
)

const (
	root             = "ignite.graphql"
	handlerConfig    = root + ".handler"
	pretty           = handlerConfig + ".pretty"
	enableGraphiQL   = handlerConfig + ".graphiQL"
	enablePlayground = handlerConfig + ".playground"
)

func init() {
	config.Add(pretty, false, "enable/disable pretty print")
	config.Add(enableGraphiQL, false, "enable/disable GraphiQL")
	config.Add(enablePlayground, true, "enable/disable Playground")
}

// DefaultHandlerConfig unmarshals the default graphql handler config.
func DefaultHandlerConfig() (*handler.Config, error) {
	return ignite.NewOptionsWithPath[handler.Config](handlerConfig)
}
