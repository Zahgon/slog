package slog

import (
	"io"
)

// SugaredLoggerFn func type.
type SugaredLoggerFn func(sl *SugaredLogger)

// SugaredLogger Is a fast and usable Logger, which already contains
// the default formatting and handling capabilities
type SugaredLogger struct {
	*Logger
	// Formatter log message formatter. default use TextFormatter
	Formatter Formatter
	// Output writer
	Output io.Writer
	// Level for log handling. if log record level <= Level, it will be record. default: DebugLevel
	//
	// TIP: setting the level to lower will ignore more logs.
	Level Level
}

// NewStd logger instance, alias of NewStdLogger()
func NewStd(fns ...SugaredLoggerFn) *SugaredLogger { _ = "STUB: not implemented"; return nil }

// NewStdLogger instance
func NewStdLogger(fns ...SugaredLoggerFn) *SugaredLogger { _ = "STUB: not implemented"; return nil }

// sl.CallerSkip += 1

// auto enable console color

// NewSugared create new SugaredLogger. alias of NewSugaredLogger()
func NewSugared(out io.Writer, level Level, fns ...SugaredLoggerFn) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

// NewSugaredLogger create new SugaredLogger
func NewSugaredLogger(output io.Writer, level Level, fns ...SugaredLoggerFn) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

// default value

// NOTICE: use self as a log handler

// NewJSONSugared create new SugaredLogger with JSONFormatter
func NewJSONSugared(out io.Writer, level Level, fns ...SugaredLoggerFn) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

// Config current logger
func (sl *SugaredLogger) Config(fns ...SugaredLoggerFn) *SugaredLogger {
	_ = "STUB: not implemented"
	return nil
}

// Reset the logger
func (sl *SugaredLogger) Reset() { _ = "STUB: not implemented"; return }

// IsHandling Check if the current level can be handling
func (sl *SugaredLogger) IsHandling(level Level) bool { _ = "STUB: not implemented"; return false }

// Handle log record
func (sl *SugaredLogger) Handle(record *Record) error { _ = "STUB: not implemented"; return nil }

// Close all log handlers, will flush and close all handlers.
//
// IMPORTANT:
//
//	if enable async/buffer mode, please call the Close() before exit.
func (sl *SugaredLogger) Close() error { _ = "STUB: not implemented"; return nil }

// TIP: must exclude self, because self is a handler

// Flush all logs. alias of the FlushAll()
func (sl *SugaredLogger) Flush() error { _ = "STUB: not implemented"; return nil }

// FlushAll all logs
func (sl *SugaredLogger) FlushAll() error { _ = "STUB: not implemented"; return nil }
