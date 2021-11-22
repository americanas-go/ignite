package zap

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/log/contrib/go.uber.org/zap.v1"
)

// NewOptions returns configured zap logger options.
func NewOptions() (*zap.Options, error) {
	o := &zap.Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
