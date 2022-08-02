package main

import (
	"os"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/go.uber.org/zap.v1"
)

func init() {
	os.Setenv("IGNITE_ZAP_CONSOLE_FORMATTER", "JSON")
	os.Setenv("IGNITE_ZAP_CONSOLE_LEVEL", "DEBUG")
}

func main() {
	ignite.Boot()
	zap.NewLogger()
}
