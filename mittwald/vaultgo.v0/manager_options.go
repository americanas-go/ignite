package vault

import (
	"github.com/americanas-go/config"
)

type ManagerOptions struct {
	SecretPath string
	Watcher    struct {
		Enabled   bool
		Increment int
	}
	Keys map[string]string
}

func NewManagerOptionsWithPath(path string) (opts *ManagerOptions, err error) {

	opts = &ManagerOptions{}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
