package compress

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// Options compress plugin for fiber options.
type Options struct {
	Enabled bool
	Level   int
}

// GetLevel returns current compress level.
func (o *Options) GetLevel() compress.Level {
	switch config.Int(level) {
	case -1:
		return compress.LevelDisabled
	case 1:
		return compress.LevelBestSpeed
	case 2:
		return compress.LevelBestCompression
	default:
		return compress.LevelDefault
	}
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
