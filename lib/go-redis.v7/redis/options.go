package redis

import (
	"time"
)

// Options represents redis client set options.
type Options struct {
	Password           string        `default:"" hide:"true" desc:"optional password. Must match the password specified in the requirepass server configuration option"`
	MaxRetries         int           `default:"0" desc:"maximum number of retries before giving up"`
	MinRetryBackoff    time.Duration `default:"8ms" desc:"minimum backoff between each retry"`
	MaxRetryBackoff    time.Duration `default:"512ms" desc:"maximum backoff between each retry"`
	DialTimeout        time.Duration `default:"5s" desc:"dial timeout for establishing new connections"`
	ReadTimeout        time.Duration `default:"3s" desc:"timeout for socket reads. If reached, commands will fail with a timeout instead of blocking. Use value -1 for no timeout and 0 for default"`
	WriteTimeout       time.Duration `default:"3s" desc:"timeout for socket writes. If reached, commands will fail"`
	PoolSize           int           `default:"10" desc:"maximum number of socket connections"`
	MinIdleConns       int           `default:"2" desc:"minimum number of idle connections which is useful when establishing new connection is slow"`
	MaxConnAge         time.Duration `default:"0" desc:"connection age at which client retires (closes) the connection"`
	PoolTimeout        time.Duration `default:"4s" desc:"amount of time client waits for connection if all connections are busy before returning an error"`
	IdleTimeout        time.Duration `default:"5m" desc:"amount of time after which client closes idle connections. Should be less than server's timeout"`
	IdleCheckFrequency time.Duration `default:"1m" desc:"frequency of idle checks made by idle connections reaper. Default is 1 minute. -1 disables idle connections reaper, but idle connections are still discarded by the client if idleTimeout is set"`
	Client             ClientOptions
	Cluster            ClusterOptions
	Sentinel           SentinelOptions
	Plugins            PluginsOptions
}

// options root path
func (o *Options) Root() string {
	return "ignite.redis"
}

func (o *Options) PostLoad() error {
	return nil
}

// represents the redis client options.
type ClientOptions struct {
	Addr    string `default:"127.0.0.1:6379" desc:"host:port address"`
	Network string `default:"tcp" desc:"the network type, either tcp or unix"`
	DB      int    `config:"db" default:"0" desc:"database to be selected after connecting to the server"`
}

// represents a redis cluster client options.
type ClusterOptions struct {
	Enabled        bool     `default:"false" desc:"whether is true a cluster client will be set up."`
	Addrs          []string `default:"127.0.0.1:6379" desc:"a seed list of host:port addresses of cluster nodes"`
	MaxRedirects   int      `config:"maxredirects" default:"8" desc:"the maximum number of retries before giving up"`
	ReadOnly       bool     `config:"readonly" default:"false" desc:"enables read-only commands on slave nodes"`
	RouteByLatency bool     `config:"routebylatency" default:"false" desc:"allows routing read-only commands to the closest master or slave node"`
	RouteRandomly  bool     `config:"routerandomly" default:"false" desc:"allows routing read-only commands to the random master or slave node"`
}

// represents redis sentinel options.
type SentinelOptions struct {
	MasterName string   `config:"masterName" default:"" desc:"redis sentinel master name"`
	Addrs      []string `default:"" desc:"redis sentinel addr list host:port"`
	Password   string   `default:"" desc:"redis sentinel password"`
}

// plugins options
type PluginsOptions struct {
	Datadog  DatadogOptions
	Health   HealthOptions
	Newrelic NewrelicOptions
}

//  health plugin options.
type HealthOptions struct {
	Enabled     bool   `default:"true" desc:"enable/disable health plugin"`
	Name        string `default:"Redis" desc:"health name"`
	Description string `default:"default connection" desc:"define health description"`
	Required    bool   `default:"true" desc:"enable/disable health"`
}

// newrelic plugin options
type NewrelicOptions struct {
	Enabled bool `default:"false" desc:"enables/disables newrelic plugin"`
}

// datadog plugin options
type DatadogOptions struct {
	Enabled       bool    `default:"false" desc:"enables/disables datadog plugin"`
	ServiceName   string  `default:"Redis" desc:"sets the given service name for the client."`
	AnalyticsRate float64 `default:"-1" desc:"the sampling rate for Trace Analytics events correlated to started spans. From 0 to 1."`
}
