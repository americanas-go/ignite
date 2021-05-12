package health

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-resty/resty.v2"
)

const (
	root        = resty.PluginsRoot + ".health"
	name        = root + ".name"
	host        = root + ".host"
	endpoint    = root + ".endpoint"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	config.Add(name, "rest api", "health name")
	config.Add(host, "", "health host")
	config.Add(endpoint, "/resource-status", "health host")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
