package vault

import (
	"github.com/americanas-go/ignite"
)

// ManagerOptions represents a vault client options.
type ManagerOptions struct {
	SecretPath string
	Watcher    struct {
		Enabled   bool
		Increment int
	}
	Keys map[string]string
}

// NewManagerOptionsWithPath unmarshals manager options based a given key path.
func NewManagerOptionsWithPath(path string) (opts *ManagerOptions, err error) {
	return ignite.NewOptionsWithPath[ManagerOptions](path)
}
