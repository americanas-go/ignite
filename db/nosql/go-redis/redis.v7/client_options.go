package redis

// ClientOptions represents the redis client options.
type ClientOptions struct {
	Addr    string
	Network string
	DB      int `config:"db"`
}
