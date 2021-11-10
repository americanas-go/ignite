package opentracing

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/client"
)

const (
	root    = client.PluginsRoot + ".opentracing"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable opentracing")
}
