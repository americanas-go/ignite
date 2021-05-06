package redis

type SentinelOptions struct {
	MasterName string `config:"masterName"`
	Addrs      []string
	Password   string
}
