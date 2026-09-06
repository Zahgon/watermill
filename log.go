package watermill

import (
	"io"
	"log"
	"sync"
	"time"
)

type LogFields map[string]interface{}

func (l LogFields) Add(newFields LogFields) LogFields {
	_ = "STUB: not implemented"
	return *new(LogFields)
}

func (l LogFields) Copy() LogFields { _ = "STUB: not implemented"; return *new(LogFields) }

type LoggerAdapter interface {
	Error(msg string, err error, fields LogFields)
	Info(msg string, fields LogFields)
	Debug(msg string, fields LogFields)
	Trace(msg string, fields LogFields)
	With(fields LogFields) LoggerAdapter
}

type NopLogger struct{}

func (NopLogger) Error(msg string, err error, fields LogFields) { _ = "STUB: not implemented"; return }
func (NopLogger) Info(msg string, fields LogFields)             { _ = "STUB: not implemented"; return }
func (NopLogger) Debug(msg string, fields LogFields)            { _ = "STUB: not implemented"; return }
func (NopLogger) Trace(msg string, fields LogFields)            { _ = "STUB: not implemented"; return }
func (l NopLogger) With(fields LogFields) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

type StdLoggerAdapter struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	TraceLogger *log.Logger

	fields LogFields
}

func NewStdLogger(debug, trace bool) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

func NewStdLoggerWithOut(out io.Writer, debug bool, trace bool) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

func (l *StdLoggerAdapter) Error(msg string, err error, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

func (l *StdLoggerAdapter) Info(msg string, fields LogFields) { _ = "STUB: not implemented"; return }

func (l *StdLoggerAdapter) Debug(msg string, fields LogFields) { _ = "STUB: not implemented"; return }

func (l *StdLoggerAdapter) Trace(msg string, fields LogFields) { _ = "STUB: not implemented"; return }

func (l *StdLoggerAdapter) With(fields LogFields) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

func (l *StdLoggerAdapter) log(logger *log.Logger, level string, msg string, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

type LogLevel uint

const (
	TraceLogLevel LogLevel = iota + 1
	DebugLogLevel
	InfoLogLevel
	ErrorLogLevel
)

type CapturedMessage struct {
	Level  LogLevel
	Time   time.Time
	Fields LogFields
	Msg    string
	Err    error
}

func (c CapturedMessage) ContentEquals(other CapturedMessage) bool {
	_ = "STUB: not implemented"
	return false
}

type CaptureLoggerAdapter struct {
	captured map[LogLevel][]CapturedMessage
	fields   LogFields
	lock     *sync.Mutex
}

func NewCaptureLogger() *CaptureLoggerAdapter { _ = "STUB: not implemented"; return nil }

func (c *CaptureLoggerAdapter) With(fields LogFields) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

func (c *CaptureLoggerAdapter) capture(level LogLevel, msg string, err error, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

func (c *CaptureLoggerAdapter) Captured() map[LogLevel][]CapturedMessage {
	_ = "STUB: not implemented"
	return nil
}

type Logfer interface {
	Logf(format string, a ...interface{})
}

func (c *CaptureLoggerAdapter) PrintCaptured(t Logfer) { _ = "STUB: not implemented"; return }

func (c *CaptureLoggerAdapter) Has(msg CapturedMessage) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CaptureLoggerAdapter) HasError(err error) bool { _ = "STUB: not implemented"; return false }

func (c *CaptureLoggerAdapter) Error(msg string, err error, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

func (c *CaptureLoggerAdapter) Info(msg string, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

func (c *CaptureLoggerAdapter) Debug(msg string, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

func (c *CaptureLoggerAdapter) Trace(msg string, fields LogFields) {
	_ = "STUB: not implemented"
	return
}
