package echo

import "github.com/americanas-go/config"

const (
	root        = "ignite.echo"
	hideBanner  = ".hidebanner"
	tp          = ".type"
	port        = ".port"
	PluginsRoot = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+hideBanner, true, "echo hide/show banner")
	config.Add(path+port, 8080, "Server http port")
	config.Add(path+tp, "REST", "defines type for applicaton")
}
