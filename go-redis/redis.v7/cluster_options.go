package redis

type ClusterOptions struct {
	Addrs          []string
	MaxRedirects   int  `config:"maxredirects"`
	ReadOnly       bool `config:"readonly"`
	RouteByLatency bool `config:"routebylatency"`
	RouteRandomly  bool `config:"routerandomly"`
}
