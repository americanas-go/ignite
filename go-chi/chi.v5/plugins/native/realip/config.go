package realip

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".realip"
	enabled = ".enabled"
)

func init() {
	ConfigAdd(root)
}

// ConfigAdd adds config from path
func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable realip middleware")
}
