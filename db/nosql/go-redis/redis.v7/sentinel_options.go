package redis

// SentinelOptions represents redis sentinel options.
type SentinelOptions struct {
	MasterName string `config:"masterName"`
	Addrs      []string
	Password   string
}
