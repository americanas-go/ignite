package prometheus

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/rpc/google.golang.org/grpc.v1/server"
)

const (
	root    = server.PluginsRoot + ".prometheus"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable prometheus")
}
