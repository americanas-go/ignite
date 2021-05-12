package opentracing

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/google.golang.org/grpc.v1/server"
)

const (
	root    = server.PluginsRoot + ".opentracing"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable opentracing")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
