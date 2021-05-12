package health

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go.mongodb.org/mongo-driver.v1"
)

const (
	root        = mongo.PluginsRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	config.Add(name, "mongo", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
