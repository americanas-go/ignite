package redis

// ClusterOptions represents a redis cluster client options.
type ClusterOptions struct {
	Addrs          []string
	MaxRedirects   int
	ReadOnly       bool
	RouteByLatency bool
	RouteRandomly  bool
}
