package redis

type ClusterOptions struct {
	Addrs          []string
	MaxRedirects   int
	ReadOnly       bool
	RouteByLatency bool
	RouteRandomly  bool
}
