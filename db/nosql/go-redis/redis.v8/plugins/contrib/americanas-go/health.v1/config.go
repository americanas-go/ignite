package health

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/db/nosql/go-redis/redis.v8"
)

const (
	root        = redis.PluginsRoot + ".health"
	name        = ".name"
	description = ".description"
	required    = ".required"
	enabled     = ".enabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+name, "redis", "health name")
	config.Add(path+description, "default connection", "define health description")
	config.Add(path+required, true, "define health description")
	config.Add(path+enabled, true, "enable/disable health")
}
