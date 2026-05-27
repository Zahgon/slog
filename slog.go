/*
Package slog Lightweight, extensible, configurable logging library written in Go.

Source code and other details for the project are available at GitHub:

	https://github.com/gookit/slog

Quick usage:

	package main

	import (
		"github.com/gookit/slog"
	)

	func main() {
		slog.Info("info log message")
		slog.Warn("warning log message")
		slog.Infof("info log %s", "message")
		slog.Debugf("debug %s", "message")
	}

More usage please see README.
*/
package slog

import (
	"context"
	"time"
)

//
// ------------------------------------------------------------
// Global std logger operate
// ------------------------------------------------------------
//

// std logger is a SugaredLogger.
// It is directly available without any additional configuration
var std = NewStdLogger()

// Std get std logger
func Std() *SugaredLogger {
	_ = "STUB: not implemented"

	// Reset the std logger and reset exit handlers
	return nil
}

func Reset() { _ = "STUB: not implemented"; return }

// new std

// Configure the std logger
func Configure(fn func(l *SugaredLogger)) {
	_ = "STUB: not implemented"

	// SetExitFunc to the std logger
	return
}

func SetExitFunc(fn func(code int)) {
	_ = "STUB: not implemented"

	// Exit runs all exit handlers and then terminates the program using os.Exit(code)
	return
}

func Exit(code int) {
	_ = "STUB: not implemented"

	// Close logger, flush and close all handlers.
	//
	// IMPORTANT: please call Close() before app exit.
	return
}

func Close() error {
	_ = "STUB: not implemented"

	// MustClose logger, flush and close all handlers.
	//
	// IMPORTANT: please call Close() before app exit.
	return nil
}

func MustClose() { _ = "STUB: not implemented"; return }

// Flush log messages
func Flush() error {
	_ = "STUB: not implemented"

	// MustFlush log messages
	return nil
}

func MustFlush() { _ = "STUB: not implemented"; return }

// FlushTimeout flush logs with timeout.
func FlushTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

// FlushDaemon run flush handle on daemon.
//
// Usage please see slog_test.ExampleFlushDaemon()
func FlushDaemon(onStops ...func()) { _ = "STUB: not implemented"; return }

// StopDaemon stop flush daemon
func StopDaemon() {
	_ = "STUB: not implemented"

	// SetLogLevel max level for the std logger
	return
}

func SetLogLevel(l Level) {
	_ = "STUB: not implemented"

	// SetLevelByName set max log level by name. eg: "info", "debug" ...
	return
}

func SetLevelByName(name string) { _ = "STUB: not implemented"; return }

// SetFormatter to std logger
func SetFormatter(f Formatter) {
	_ = "STUB: not implemented"

	// GetFormatter of the std logger
	return
}

func GetFormatter() Formatter {
	_ = "STUB: not implemented"
	return *

	// AddHandler to the std logger
	new(Formatter)
}

func AddHandler(h Handler) {
	_ = "STUB: not implemented"

	// PushHandler to the std logger
	return
}

func PushHandler(h Handler) {
	_ = "STUB: not implemented"

	// AddHandlers to the std logger
	return
}

func AddHandlers(hs ...Handler) { _ = "STUB: not implemented"; return }

// PushHandlers to the std logger
func PushHandlers(hs ...Handler) { _ = "STUB: not implemented"; return }

// AddProcessor to the logger
func AddProcessor(p Processor) {
	_ = "STUB: not implemented"

	// AddProcessors to the logger
	return
}

func AddProcessors(ps ...Processor) { _ = "STUB: not implemented"; return }

// -------------------------- New sub-logger -----------------------------

// NewSub returns a new SubLogger on the std logger.
func NewSub() *SubLogger { _ = "STUB: not implemented"; return nil }

// -------------------------- New record with log data, fields -----------------------------

// WithExtra new record with extra data
func WithExtra(ext M) *Record { _ = "STUB: not implemented"; return nil }

// WithData new record with data
func WithData(data M) *Record { _ = "STUB: not implemented"; return nil }

// WithValue new record with data value
func WithValue(key string, value any) *Record { _ = "STUB: not implemented"; return nil }

// WithField new record with field.
//
// **NOTE**: add field need config Formatter template fields.
func WithField(name string, value any) *Record { _ = "STUB: not implemented"; return nil }

// WithFields new record with fields
//
// **NOTE**: add field need config Formatter template fields.
func WithFields(fields M) *Record { _ = "STUB: not implemented"; return nil }

// WithContext new record with context
func WithContext(ctx context.Context) *Record { _ = "STUB: not implemented"; return nil }

// region Add log messages
// -------------------------- Add log messages with level -----------------------------

// Log logs a message with level
func Log(level Level, args ...any) { _ = "STUB: not implemented"; return }

// Print logs a message at level PrintLevel
func Print(args ...any) { _ = "STUB: not implemented"; return }

// Println logs a message at level PrintLevel
func Println(args ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message at level PrintLevel
func Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Trace logs a message at level TraceLevel
func Trace(args ...any) { _ = "STUB: not implemented"; return }

// Tracef logs a message at level TraceLevel
func Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

// TraceCtx logs a message at level TraceLevel with context
func TraceCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// TracefCtx logs a message at level TraceLevel with context
func TracefCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// Debug logs a message at level DebugLevel
func Debug(args ...any) { _ = "STUB: not implemented"; return }

// Debugf logs a message at level DebugLevel
func Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

// DebugCtx logs a message at level DebugLevel with context
func DebugCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// DebugfCtx logs a message at level DebugLevel with context
func DebugfCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// Info logs a message at level InfoLevel
func Info(args ...any) { _ = "STUB: not implemented"; return }

// Infof logs a message at level InfoLevel
func Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

// InfoCtx logs a message at level InfoLevel with context
func InfoCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// InfofCtx logs a message at level InfoLevel with context
func InfofCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// Notice logs a message at level NoticeLevel
func Notice(args ...any) { _ = "STUB: not implemented"; return }

// Noticef logs a message at level NoticeLevel
func Noticef(format string, args ...any) { _ = "STUB: not implemented"; return }

// NoticeCtx logs a message at level NoticeLevel with context
func NoticeCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// NoticefCtx logs a message at level NoticeLevel with context
func NoticefCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// Warn logs a message at level WarnLevel
func Warn(args ...any) { _ = "STUB: not implemented"; return }

// Warnf logs a message at level WarnLevel
func Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

// WarnCtx logs a message at level Warn with a context
func WarnCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// WarnfCtx logs a message at level Warn with a context
func WarnfCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// Error logs a message at level Error
func Error(args ...any) { _ = "STUB: not implemented"; return }

// Errorf logs a message at level Error
func Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

// ErrorT logs a error type at level Error
func ErrorT(err error) { _ = "STUB: not implemented"; return }

// ErrorCtx logs a message at level Error with context
func ErrorCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// ErrorfCtx logs a message at level Error with context
func ErrorfCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// EStack logs a error message and with call stack.
// func EStack(args ...any) {
// 	std.WithExtra(map[string]any{"stack": goinfo.GetCallerInfo(2)}).
// 		log(ErrorLevel, args)
// }

// Fatal logs a message at level Fatal
func Fatal(args ...any) { _ = "STUB: not implemented"; return }

// Fatalf logs a message at level Fatal
func Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// FatalErr logs a message at level Fatal on err is not nil
func FatalErr(err error) { _ = "STUB: not implemented"; return }

// FatalCtx logs a message at level Fatal with context
func FatalCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// FatalfCtx logs a message at level Fatal with context
func FatalfCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }

// Panic logs a message at level Panic
func Panic(args ...any) { _ = "STUB: not implemented"; return }

// Panicf logs a message at level Panic
func Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }

// PanicErr logs a message at level Panic on err is not nil
func PanicErr(err error) { _ = "STUB: not implemented"; return }

// PanicCtx logs a message at level panic with context
func PanicCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// PanicfCtx logs a message at level panic with context
func PanicfCtx(ctx context.Context, format string, args ...any) { _ = "STUB: not implemented"; return }
