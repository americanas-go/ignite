package echo

import (
	"fmt"
	"io"

	"github.com/americanas-go/log"
	l "github.com/labstack/gommon/log"
)

// WrappedLogger represents a wrapped logger with util methods.
type WrappedLogger struct {
	logger log.Logger
}

// Output returns logger output writer.
func (wl WrappedLogger) Output() io.Writer {
	return wl.logger.Output()
}

// Prefix not implemented yet.
func (wl WrappedLogger) Prefix() string {
	wl.logger.Errorf("Prefix(): implement me")
	return ""
}

// SetPrefix not implemented yet.
func (wl WrappedLogger) SetPrefix(p string) {
	wl.logger.Errorf("WrappedLogger.SetPrefix(p string): implement me")
}

// Level not implemented yet.
func (wl WrappedLogger) Level() l.Lvl {
	wl.logger.Errorf("WrappedLogger.Level(): implement me")
	return l.INFO
}

//SetLevel not implemented yet.
func (wl WrappedLogger) SetLevel(v l.Lvl) {
	wl.logger.Errorf("WrappedLogger.SetLevel(v log.Lvl): implement me")
}

// SetHeader not implemented yet.
func (wl WrappedLogger) SetHeader(h string) {
	wl.logger.Errorf("WrappedLogger.SetHeader(h string): implement me")
}

// Level not implemented yet.
func (wl WrappedLogger) Printj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Printj(j l.JSON): implement me")
}

// Debugj not implemented yet.
func (wl WrappedLogger) Debugj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Debugj(j l.JSON) implement me")
}

// Infoj not implemented yet.
func (wl WrappedLogger) Infoj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Infoj(j l.JSON): implement me")
}

// Warnj not implemented yet.
func (wl WrappedLogger) Warnj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Warnj(j l.JSON): implement me")
}

// Warnj not implemented yet.
func (wl WrappedLogger) Errorj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Errorj(j l.JSON): implement me")
}

// Warnj not implemented yet.
func (wl WrappedLogger) Fatalj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Fatalj(j l.JSON): implement me")
}

// Warnj not implemented yet.
func (wl WrappedLogger) Panic(i ...interface{}) {
	wl.logger.Errorf("WrappedLogger.Panic(i ...interface{}): implement me")
}

// Warnj not implemented yet.
func (wl WrappedLogger) Panicj(j l.JSON) {
	wl.logger.Errorf("WrappedLogger.Panicj(j l.JSON): implement me")
}

// Warnj not implemented yet.
func (wl WrappedLogger) Panicf(format string, args ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(format, args...))
}

// Warnj not implemented yet.
func (wl WrappedLogger) SetOutput(w io.Writer) {
	wl.logger.Errorf("WrappedLogger.SetOutput(w io.Writer): implement me")
}

// Print logs a message.
func (wl WrappedLogger) Print(i ...interface{}) {
	wl.logger.Printf("%v", i...)
}

// Print logs a formatted message.
func (wl WrappedLogger) Printf(s string, i ...interface{}) {
	wl.logger.Printf(s, i...)
}

// Debug logs a debug message.
func (wl WrappedLogger) Debug(i ...interface{}) {
	wl.logger.Debugf(fmt.Sprint(i...))
}

// Debugf logs a formatted debug message.
func (wl WrappedLogger) Debugf(s string, i ...interface{}) {
	wl.logger.Debugf(fmt.Sprintf(s, i...))
}

// Warnf logs an info message.
func (wl WrappedLogger) Info(i ...interface{}) {
	wl.logger.Infof(fmt.Sprint(i...))
}

// Infof logs a formatted info message.
func (wl WrappedLogger) Infof(s string, i ...interface{}) {
	wl.logger.Infof(fmt.Sprintf(s, i...))
}

// Warn logs a warn message.
func (wl WrappedLogger) Warn(i ...interface{}) {
	wl.logger.Warnf(fmt.Sprint(i...))
}

// Warnf logs a formatted warn message.
func (wl WrappedLogger) Warnf(s string, i ...interface{}) {
	wl.logger.Warnf(fmt.Sprintf(s, i...))
}

// Error logs an error message.
func (wl WrappedLogger) Error(i ...interface{}) {
	wl.logger.Errorf(fmt.Sprint(i...))
}

// Errorf logs a formatted error message.
func (wl WrappedLogger) Errorf(s string, i ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(s, i...))
}

// Fatal logs a fatal message.
func (wl WrappedLogger) Fatal(i ...interface{}) {
	wl.logger.Fatal(fmt.Sprint(i...))
}

// Fatalf logs a formatted fatal message.
func (wl WrappedLogger) Fatalf(s string, i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprintf(s, i...))
}

// WrapLogger wraps an instace of logger interface.
func WrapLogger(l log.Logger) *WrappedLogger {
	return &WrappedLogger{logger: l}
}
