package godror

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root            = "ignite.godror"
	dataSourceName  = ".dataSourceName"
	connMaxLifetime = ".connMaxLifetime"
	maxIdleConns    = ".maxIdleConns"
	maxOpenConns    = ".maxOpenConns"
	PluginsRoot     = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+dataSourceName, "", "database name and connection information")
	config.Add(path+connMaxLifetime, 0*time.Second, "sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever")
	config.Add(path+maxIdleConns, 2, "sets the maximum number of connections in the idle connection pool.")
	config.Add(path+maxOpenConns, 5, "sets the maximum number of open connections to the database.")
}
