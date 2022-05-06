package redis

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root               = "ignite.redis"
	pp                 = ".password"
	maxRetries         = ".maxRetries"
	minRetryBackoff    = ".minRetryBackoff"
	maxRetryBackoff    = ".maxRetryBackoff"
	dialTimeout        = ".dialTimeout"
	readTimeout        = ".readTimeout"
	writeTimeout       = ".writeTimeout"
	poolSize           = ".poolSize"
	minIdleConns       = ".minIdleConns"
	maxConnAge         = ".maxConnAge"
	poolTimeout        = ".poolTimeout"
	idleTimeout        = ".idleTimeout"
	idleCheckFrequency = ".idleCheckFrequency"
	addr               = ".client.addr"
	network            = ".client.network"
	db                 = ".client.db"
	sentinelMaster     = ".sentinel.masterName"
	sentinelAddr       = ".sentinel.addrs"
	sentinelPP         = ".sentinel.password"
	addrs              = ".cluster.addrs"
	maxRedirects       = ".cluster.maxRedirects"
	readOnly           = ".cluster.readOnly"
	routeByLatency     = ".cluster.routeByLatency"
	routeRandomly      = ".cluster.routeRandomly"
	PluginsRoot        = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+addrs, []string{"127.0.0.1:6379"}, "a seed list of host:port addresses of cluster nodes")
	config.Add(path+maxRedirects, 8, "the maximum number of retries before giving up")
	config.Add(path+readOnly, false, "enables read-only commands on slave nodes")
	config.Add(path+routeByLatency, false, "allows routing read-only commands to the closest master or slave node")
	config.Add(path+routeRandomly, false, "allows routing read-only commands to the random master or slave node")
	config.Add(path+pp, "", "optional password. Must match the password specified in the requirepass server configuration option", config.WithHide())
	config.Add(path+maxRetries, 0, "maximum number of retries before giving up")
	config.Add(path+minRetryBackoff, 8*time.Millisecond, "minimum backoff between each retry")
	config.Add(path+maxRetryBackoff, 512*time.Millisecond, "maximum backoff between each retry")
	config.Add(path+dialTimeout, 5*time.Second, "dial timeout for establishing new connections")
	config.Add(path+readTimeout, 3*time.Second, "timeout for socket reads. If reached, commands will fail with a timeout instead of blocking. Use value -1 for no timeout and 0 for default")
	config.Add(path+writeTimeout, 3*time.Second, "timeout for socket writes. If reached, commands will fail")
	config.Add(path+poolSize, 10, "maximum number of socket connections")
	config.Add(path+minIdleConns, 2, "minimum number of idle connections which is useful when establishing new connection is slow")
	config.Add(path+maxConnAge, 0*time.Millisecond, "connection age at which client retires (closes) the connection")
	config.Add(path+poolTimeout, 4*time.Second, "amount of time client waits for connection if all connections are busy before returning an error")
	config.Add(path+idleTimeout, 5*time.Minute, "amount of time after which client closes idle connections. Should be less than server's timeout")
	config.Add(path+idleCheckFrequency, 1*time.Minute, "frequency of idle checks made by idle connections reaper. Default is 1 minute. -1 disables idle connections reaper, but idle connections are still discarded by the client if idleTimeout is set")
	config.Add(path+addr, "127.0.0.1:6379", "host:port address")
	config.Add(path+network, "tcp", "the network type, either tcp or unix")
	config.Add(path+db, 0, "database to be selected after connecting to the server")
	config.Add(path+sentinelMaster, "", "redis sentinel master name")
	config.Add(path+sentinelAddr, nil, "redis sentinel addr list host:port")
	config.Add(path+sentinelPP, "", "redis sentinel password")
}
