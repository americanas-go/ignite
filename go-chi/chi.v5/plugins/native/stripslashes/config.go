package stripslashes

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".stripSlashes"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

// ConfigAdd adds config from path
func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable stripSlashes middleware")
}
