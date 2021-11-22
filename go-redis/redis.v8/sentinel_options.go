package redis

// SentinelOptions represents redis sentinel options.
type SentinelOptions struct {
	MasterName string
	Addrs      []string
	Password   string
}
