package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/google.golang.org/grpc.v1/server"
)

const (
	root    = server.PluginsRoot + ".log"
	enabled = ".enabled"
	level   = ".level"
)

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable logger")
	config.Add(path+level, "INFO", "sets log level INFO/DEBUG/TRACE")
}
