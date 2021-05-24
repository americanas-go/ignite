package redis

type SentinelOptions struct {
	MasterName string
	Addrs      []string
	Password   string
}
