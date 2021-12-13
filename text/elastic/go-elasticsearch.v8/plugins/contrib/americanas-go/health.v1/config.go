package health

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/text/elastic/go-elasticsearch.v8"
)

const (
	root        = elasticsearch.PluginsRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+name, "elasticsearch", "health name")
	config.Add(path+description, "default connection", "define health description")
	config.Add(path+required, true, "define health description")
	config.Add(path+enabled, true, "enable/disable health")
}
