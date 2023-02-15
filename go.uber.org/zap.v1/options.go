package zap

import (
	"github.com/americanas-go/ignite"
	"github.com/americanas-go/log/contrib/go.uber.org/zap.v1"
)

// NewOptions returns configured zap logger options.
func NewOptions() (*zap.Options, error) {
	return ignite.NewOptionsWithPath[zap.Options](root)
}
