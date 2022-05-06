package godror

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root                = "ignite.godror"
	connectString       = ".connectString"
	pu                  = ".username"
	pp                  = ".password"
	maxLifetime         = ".maxLifetime"
	sessionTimeout      = ".sessionTimeout"
	waitTimeout         = ".waitTimeout"
	maxSessions         = ".maxSessions"
	minSessions         = ".minSessions"
	maxSessionsPerShard = ".maxSessionsPerShard"
	sessionIncrement    = ".sessionIncrement"
	PluginsRoot         = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+connectString, "localhost:1521/database?connect_timeout=2", "sets database connection string")
	config.Add(path+pu, "", "sets database username")
	config.Add(path+pp, "", "sets database password")
	config.Add(path+maxLifetime, 1*time.Hour, "sets the maximum amount of time a connection may be reused. If d <= 0, connections are reused forever")
	config.Add(path+sessionTimeout, 5*time.Minute, "sets the session timeout")
	config.Add(path+waitTimeout, 30*time.Second, "sets the wait timeout")
	config.Add(path+maxSessions, 1000, "sets the maximum sessions to the database")
	config.Add(path+minSessions, 1, "sets the minimum sessions to the database")
	config.Add(path+sessionIncrement, 1, "sets the session increment")
	config.Add(path+maxSessionsPerShard, 0, "sets the max sessions per shard")
}
