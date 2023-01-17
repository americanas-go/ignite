package hystrix

import (
	"strings"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/americanas-go/grapper.v1"
)

const (
	root    = grapper.PluginsRoot + ".hystrix"
	enabled = ".enabled"
)

func ConfigAdd(name string) {
	path := strings.Join([]string{root, ".", name}, "")
	config.Add(path+enabled, true, "enable/disable hystrix grapper middleware")
}
