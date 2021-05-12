package health

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/elastic/go-elasticsearch.v8"
)

const (
	root        = elasticsearch.PluginsRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	config.Add(name, "elasticsearch", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
