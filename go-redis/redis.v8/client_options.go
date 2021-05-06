package redis

type ClientOptions struct {
	Addr    string
	Network string
	DB      int `config:"db"`
}
