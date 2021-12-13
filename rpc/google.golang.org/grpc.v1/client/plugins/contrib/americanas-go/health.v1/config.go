package health

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/rpc/google.golang.org/grpc.v1/client"
)

const (
	root        = client.PluginsRoot + ".health"
	name        = ".name"
	description = ".description"
	required    = ".required"
	enabled     = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+name, "grpc", "health name")
	config.Add(path+description, "default connection", "define health description")
	config.Add(path+required, true, "define health description")
	config.Add(path+enabled, true, "enable/disable health")
}
