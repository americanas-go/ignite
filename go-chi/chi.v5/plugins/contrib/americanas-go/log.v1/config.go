package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".logger"
	enabled = ".enabled"
	level   = ".level"
)

func init() {
	ConfigAdd(root)
}

// ConfigAdd adds config from path
func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable logger middleware")
	config.Add(path+level, "INFO", "sets log level INFO/DEBUG/TRACE")
}
