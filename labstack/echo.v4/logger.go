package echo

import (
	"fmt"
	"io"

	"github.com/americanas-go/log"
	l "github.com/labstack/gommon/log"
)

type wrappedLogger struct {
	logger log.Logger
}

func (wl wrappedLogger) Output() io.Writer {
	return wl.logger.Output()
}

func (wl wrappedLogger) Prefix() string {
	wl.logger.Errorf("Prefix(): implement me")
	return ""
}

func (wl wrappedLogger) SetPrefix(p string) {
	wl.logger.Errorf("wrappedLogger.SetPrefix(p string): implement me")
}

func (wl wrappedLogger) Level() l.Lvl {
	wl.logger.Errorf("wrappedLogger.Level(): implement me")
	return l.INFO
}

func (wl wrappedLogger) SetLevel(v l.Lvl) {
	wl.logger.Errorf("wrappedLogger.SetLevel(v log.Lvl): implement me")
}

func (wl wrappedLogger) SetHeader(h string) {
	wl.logger.Errorf("wrappedLogger.SetHeader(h string): implement me")
}

func (wl wrappedLogger) Printj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Printj(j l.JSON): implement me")
}

func (wl wrappedLogger) Debugj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Debugj(j l.JSON) implement me")
}

func (wl wrappedLogger) Infoj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Infoj(j l.JSON): implement me")
}

func (wl wrappedLogger) Warnj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Warnj(j l.JSON): implement me")
}

func (wl wrappedLogger) Errorj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Errorj(j l.JSON): implement me")
}

func (wl wrappedLogger) Fatalj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Fatalj(j l.JSON): implement me")
}

func (wl wrappedLogger) Panic(i ...interface{}) {
	wl.logger.Errorf("wrappedLogger.Panic(i ...interface{}): implement me")
}

func (wl wrappedLogger) Panicj(j l.JSON) {
	wl.logger.Errorf("wrappedLogger.Panicj(j l.JSON): implement me")
}

func (wl wrappedLogger) Panicf(format string, args ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(format, args...))
}

func (wl wrappedLogger) SetOutput(w io.Writer) {
	wl.logger.Errorf("wrappedLogger.SetOutput(w io.Writer): implement me")
}

func (wl wrappedLogger) Print(i ...interface{}) {
	wl.logger.Printf("%v", i...)
}

func (wl wrappedLogger) Printf(s string, i ...interface{}) {
	wl.logger.Printf(s, i...)
}

func (wl wrappedLogger) Debug(i ...interface{}) {
	wl.logger.Debugf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Debugf(s string, i ...interface{}) {
	wl.logger.Debugf(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Info(i ...interface{}) {
	wl.logger.Infof(fmt.Sprint(i...))
}

func (wl wrappedLogger) Infof(s string, i ...interface{}) {
	wl.logger.Infof(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Warn(i ...interface{}) {
	wl.logger.Warnf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Warnf(s string, i ...interface{}) {
	wl.logger.Warnf(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Error(i ...interface{}) {
	wl.logger.Errorf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Errorf(s string, i ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Fatal(i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Fatalf(s string, i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprintf(s, i...))
}

func WrapLogger(l log.Logger) wrappedLogger {
	return wrappedLogger{logger: l}
}
