package datadog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/rpc/google.golang.org/grpc.v1/client"
)

const (
	root    = client.PluginsRoot + ".datadog"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable datadog")
}
