package promotheus

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/afex/hystrix-go.v0"
)

const (
	root      = hystrix.PluginsRoot + ".prometheus"
	namespace = root + ".namespace"
	labels    = root + ".labels"
	enabled   = root + ".enabled"
)

func init() {
	config.Add(namespace, "hystrix", "defines hystrix prometheus namespace")
	config.Add(labels, map[string]interface{}{}, "defines hystrix prometheus labels")
	config.Add(enabled, true, "enabled/disable hystrix prometheus")
}
