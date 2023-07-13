package elasticsearch

import "github.com/americanas-go/log"

type DebugLogger struct {
}

func (l *DebugLogger) Printf(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}
