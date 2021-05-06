package mongo

import "github.com/americanas-go/config"

const (
	root        = "ignite.mongo"
	uri         = root + ".uri"
	authRoot    = root + ".auth"
	username    = authRoot + ".username"
	password    = authRoot + ".password"
	PluginsRoot = root + ".plugins"
)

func init() {
	config.Add(uri, "mongodb://localhost:27017/temp", "define mongodb uri")
	config.Add(username, "", "define mongodb username", config.WithHide())
	config.Add(password, "", "define mongodb password", config.WithHide())
}
