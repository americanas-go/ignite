package compress

import (
	"github.com/americanas-go/config"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Options struct {
	Enabled bool
	Level   int
}

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

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string) (opts *Options, err error) {
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
