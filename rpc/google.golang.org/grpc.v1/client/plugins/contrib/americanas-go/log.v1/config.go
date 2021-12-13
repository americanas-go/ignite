package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/rpc/google.golang.org/grpc.v1/client"
)

const (
	root    = client.PluginsRoot + ".log"
	enabled = ".enabled"
	level   = ".level"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable logger")
	config.Add(path+level, "INFO", "sets log level INFO/DEBUG/TRACE")
}
