package log

import "github.com/americanas-go/config"

const (
	root = "ignite.log"
	impl = root + ".impl"
)

func init() {
	config.Add(impl, "LOGRUS", "defines log implementation LOGRUS/ZAP/ZEROLOG")
}

func Impl() string {
	return config.String(impl)
}
