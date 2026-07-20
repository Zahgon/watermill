package watermill

import (
	"log/slog"
)

const LevelTrace = slog.LevelDebug - 4

func slogAttrsFromFields(fields LogFields) []any { _ = "STUB: not implemented"; return nil }

type SlogLoggerAdapter struct {
	slog *slog.Logger

	watermillLevelToSlog map[slog.Level]slog.Level
}

func (s *SlogLoggerAdapter) Error(msg string, err error, fields LogFields) {
	_ = "STUB: not implemented"
	return
}

func (s *SlogLoggerAdapter) Info(msg string, fields LogFields) { _ = "STUB: not implemented"; return }

func (s *SlogLoggerAdapter) Debug(msg string, fields LogFields) { _ = "STUB: not implemented"; return }

func (s *SlogLoggerAdapter) Trace(msg string, fields LogFields) { _ = "STUB: not implemented"; return }

func (s *SlogLoggerAdapter) log(level slog.Level, msg string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (s *SlogLoggerAdapter) With(fields LogFields) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

func NewSlogLogger(logger *slog.Logger) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}

func NewSlogLoggerWithLevelMapping(logger *slog.Logger, watermillLevelToSlog map[slog.Level]slog.Level) LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(LoggerAdapter)
}
