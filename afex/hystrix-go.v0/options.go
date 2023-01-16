package hystrix

import (
	"strings"

	"github.com/americanas-go/ignite"
)

type Options struct {
	Enabled                bool
	Timeout                int
	RequestVolumeThreshold int
	ErrorPercentThreshold  int
	MaxConcurrentRequests  int
	SleepWindow            int
}

// NewOptionsFromCommand unmarshals options based a given key path.
func NewOptionsFromCommand(cmd string) (*Options, error) {
	path := strings.Join([]string{cmdRoot, cmd}, ".")
	return ignite.NewOptionsWithPath[Options](path)
}
