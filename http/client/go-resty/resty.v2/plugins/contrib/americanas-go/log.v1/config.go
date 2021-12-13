package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/client/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".log"
	enabled = ".enabled"
	level   = ".level"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable logger")
	config.Add(path+level, "DEBUG", "sets log level INFO/DEBUG/TRACE")
}
