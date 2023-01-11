package redis

import (
	"github.com/americanas-go/cache/contrib/go-redis/redis.v8"
	"github.com/americanas-go/config"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*redis.Options, error) {
	o := &redis.Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *redis.Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
