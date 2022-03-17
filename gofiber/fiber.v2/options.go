package fiber

import (
	"github.com/gofiber/fiber/v2"
)

// Options represents fiber server options
type Options struct {
	Port   int
	Type   string
	Config *fiber.Config
}
