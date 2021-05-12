package realip

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
)

const (
	enabled = chi.PluginsRoot + ".realip.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable realip middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
