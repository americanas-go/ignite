package compress

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/gofiber/fiber.v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

const (
	root    = fiber.PluginsRoot + ".compress"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable compress middleware")
	config.Add(level, 0, "compress level (disabled: -1, default: 0, best speed: 1, best compression: 2)")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func Level() compress.Level {
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
