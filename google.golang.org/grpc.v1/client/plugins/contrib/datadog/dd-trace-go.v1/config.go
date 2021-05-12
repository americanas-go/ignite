package datadog

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/google.golang.org/grpc.v1/client"
)

const (
	root    = client.PluginsRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
