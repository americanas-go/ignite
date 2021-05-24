package echo

import (
	"fmt"
	"io"

	"github.com/americanas-go/log"
	l "github.com/labstack/gommon/log"
)

type WrappedLogger struct {
	logger log.Logger
}

func (wl WrappedLogger) Output() io.Writer {
	return wl.logger.Output()
}

func (wl WrappedLogger) Prefix() string {
	wl.logger.Errorf("Prefix(): implement me")
	return ""
}

func (wl WrappedLogger) SetPrefix(p string) {
	wl.logger.Errorf("WrappedLogger.SetPrefix(p string): implement me")
}

func (wl WrappedLogger) Level() l.Lvl {
	wl.logger.Errorf("WrappedLogger.Level(): implement me")
	return l.INFO
}

func (wl WrappedLogger) SetLevel(v l.Lvl) {
	wl.logger.Errorf("WrappedLogger.SetLevel(v log.Lvl): implement me")
}

func (wl WrappedLogger) SetHeader(h string) {
	wl.logger.Errorf("WrappedLogger.SetHeader(h string): implement me")
}

func (wl WrappedLogger) Printj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Printj(j l.JSON): implement me")
}

func (wl WrappedLogger) Debugj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Debugj(j l.JSON) implement me")
}

func (wl WrappedLogger) Infoj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Infoj(j l.JSON): implement me")
}

func (wl WrappedLogger) Warnj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Warnj(j l.JSON): implement me")
}

func (wl WrappedLogger) Errorj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Errorj(j l.JSON): implement me")
}

func (wl WrappedLogger) Fatalj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Fatalj(j l.JSON): implement me")
}

func (wl WrappedLogger) Panic(i ...interface{}) {
	wl.logger.Errorf("WrappedLogger.Panic(i ...interface{}): implement me")
}

func (wl WrappedLogger) Panicj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Panicj(j l.JSON): implement me")
}

func (wl WrappedLogger) Panicf(format string, args ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(format, args...))
}

func (wl WrappedLogger) SetOutput(w io.Writer) {
	wl.logger.Errorf("WrappedLogger.SetOutput(w io.Writer): implement me")
}

func (wl WrappedLogger) Print(i ...interface{}) {
	wl.logger.Printf("%v", i...)
}

func (wl WrappedLogger) Printf(s string, i ...interface{}) {
	wl.logger.Printf(s, i...)
}

func (wl WrappedLogger) Debug(i ...interface{}) {
	wl.logger.Debugf(fmt.Sprint(i...))
}

func (wl WrappedLogger) Debugf(s string, i ...interface{}) {
	wl.logger.Debugf(fmt.Sprintf(s, i...))
}

func (wl WrappedLogger) Info(i ...interface{}) {
	wl.logger.Infof(fmt.Sprint(i...))
}

func (wl WrappedLogger) Infof(s string, i ...interface{}) {
	wl.logger.Infof(fmt.Sprintf(s, i...))
}

func (wl WrappedLogger) Warn(i ...interface{}) {
	wl.logger.Warnf(fmt.Sprint(i...))
}

func (wl WrappedLogger) Warnf(s string, i ...interface{}) {
	wl.logger.Warnf(fmt.Sprintf(s, i...))
}

func (wl WrappedLogger) Error(i ...interface{}) {
	wl.logger.Errorf(fmt.Sprint(i...))
}

func (wl WrappedLogger) Errorf(s string, i ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(s, i...))
}

func (wl WrappedLogger) Fatal(i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprint(i...))
}

func (wl WrappedLogger) Fatalf(s string, i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprintf(s, i...))
}

func WrapLogger(l log.Logger) *WrappedLogger {
	return &WrappedLogger{logger: l}
}
