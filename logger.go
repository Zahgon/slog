package slog

import (
	"context"
	"sync"
	"time"
)

// Logger log dispatcher definition.
//
// The logger implements the `github.com/gookit/gsr.Logger`
type Logger struct {
	name string
	// lock for writing logs
	mu sync.Mutex
	// logger latest error
	err error
	// mark logger is closed
	closed bool

	// log handlers for logger
	handlers   []Handler
	processors []Processor

	// reusable empty record
	recordPool sync.Pool
	// handlers on exit.
	exitHandlers []func()
	quitDaemon   chan struct{}

	//
	// logger options
	//

	// ChannelName log channel name, default is DefaultChannelName
	ChannelName string
	// FlushInterval flush interval time. default is defaultFlushInterval=30s
	FlushInterval time.Duration
	// LowerLevelName use lower level name
	LowerLevelName bool
	// ReportCaller on writing log record
	ReportCaller bool
	CallerSkip   int
	// CallerFlag used to set caller traceback information in different modes
	CallerFlag CallerFlagMode
	// BackupArgs backup log input args to Record.Args
	BackupArgs bool
	// GlobalFields global fields. will be added to all log records
	//
	// NOTE: add field need config Formatter template fields.
	GlobalFields map[string]any
	// TimeClock custom time clock, timezone
	TimeClock ClockFn
	// custom exit, panic handler.
	ExitFunc  func(code int)
	PanicFunc func(v any)
}

// New create a new logger
func New(fns ...LoggerFn) *Logger { _ = "STUB: not implemented"; return nil }

// NewWithHandlers create a new logger with handlers
func NewWithHandlers(hs ...Handler) *Logger { _ = "STUB: not implemented"; return nil }

// NewWithConfig create a new logger with config func
func NewWithConfig(fns ...LoggerFn) *Logger { _ = "STUB: not implemented"; return nil }

// NewWithName create a new logger with name
func NewWithName(name string, fns ...LoggerFn) *Logger { _ = "STUB: not implemented"; return nil }

// exit handle
// ExitFunc:  os.Exit,

// options

// flush interval time

// NewRecord get new logger record
func (l *Logger) newRecord() *Record { _ = "STUB: not implemented"; return nil }

func (l *Logger) releaseRecord(r *Record) {
	_ = "STUB: not implemented"
	// must reset for each record
	return
}

// reuse=true: will not be released

// reset ctx data

// reset flags

//
// ---------------------------------------------------------------------------
// region Configure logger
// ---------------------------------------------------------------------------
//

// Config current logger
func (l *Logger) Config(fns ...LoggerFn) *Logger { _ = "STUB: not implemented"; return nil }

// Configure current logger. alias of Config()
func (l *Logger) Configure(fn LoggerFn) *Logger {
	_ = "STUB: not implemented"

	// RegisterExitHandler register an exit-handler on global exitHandlers
	return nil
}

func (l *Logger) RegisterExitHandler(handler func()) { _ = "STUB: not implemented"; return }

// PrependExitHandler prepend register an exit-handler on global exitHandlers
func (l *Logger) PrependExitHandler(handler func()) { _ = "STUB: not implemented"; return }

// ResetExitHandlers reset logger exitHandlers
func (l *Logger) ResetExitHandlers() { _ = "STUB: not implemented"; return }

// ExitHandlers get all exitHandlers of the logger
func (l *Logger) ExitHandlers() []func() { _ = "STUB: not implemented"; return nil }

// SetName for logger
func (l *Logger) SetName(name string) {
	_ = "STUB: not implemented"

	// Name of the logger
	return
}

func (l *Logger) Name() string {
	_ = "STUB: not implemented"

	// ---------------------------------------------------------------------------
	// region Management logger
	// ---------------------------------------------------------------------------
	return ""
}

const defaultFlushInterval = 30 * time.Second

// FlushDaemon run flush handle on daemon
//
// Usage, please refer to the FlushDaemon() on package.
func (l *Logger) FlushDaemon(onStops ...func()) { _ = "STUB: not implemented"; return }

// create a ticker

// StopDaemon stop flush daemon
func (l *Logger) StopDaemon() { _ = "STUB: not implemented"; return }

// FlushTimeout flush logs on limit time.
//
// refer from glog package
func (l *Logger) FlushTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

// Sync flushes buffered logs (if any). alias of the Flush()
func (l *Logger) Sync() error {
	_ = "STUB: not implemented"

	// Flush flushes all the logs and attempts to "sync" their data to disk.
	// l.mu is held.
	return nil
}

func (l *Logger) Flush() error { _ = "STUB: not implemented"; return nil }

// MustFlush flush logs. will panic on error
func (l *Logger) MustFlush() { _ = "STUB: not implemented"; return }

// FlushAll flushes all the logs and attempts to "sync" their data to disk.
//
// alias of the Flush()
func (l *Logger) FlushAll() error { _ = "STUB: not implemented"; return nil }

// lockAndFlushAll is like flushAll but locks l.mu first.
func (l *Logger) lockAndFlushAll() error { _ = "STUB: not implemented"; return nil }

// flush all without lock
func (l *Logger) flushAll() {
	_ = "STUB: not implemented"
	// flush from fatal down, in case there's trouble flushing.
	return
}

// MustClose close logger. will panic on error
func (l *Logger) MustClose() { _ = "STUB: not implemented"; return }

// Close the logger, will flush all logs and close all handlers
//
// IMPORTANT:
//
//	if enable async/buffer mode, please call the Close() before exit.
func (l *Logger) Close() error { _ = "STUB: not implemented"; return nil }

// VisitAll logger handlers
func (l *Logger) VisitAll(fn func(handler Handler) error) error {
	_ = "STUB: not implemented"
	return nil
}

// TIP: you can return nil for ignore error

// Reset the logger. will reset: handlers, processors, closed=false
func (l *Logger) Reset() { _ = "STUB: not implemented"; return }

// ResetProcessors for the logger
func (l *Logger) ResetProcessors() { _ = "STUB: not implemented"; return }

// ResetHandlers for the logger
func (l *Logger) ResetHandlers() { _ = "STUB: not implemented"; return }

// Exit logger handle
func (l *Logger) Exit(code int) {
	_ = "STUB: not implemented"

	// global exit handlers
	return
}

func (l *Logger) runExitHandlers() { _ = "STUB: not implemented"; return }

// DoNothingOnPanicFatal do nothing on panic or fatal level. TIP: useful on testing.
func (l *Logger) DoNothingOnPanicFatal() { _ = "STUB: not implemented"; return }

// HandlersNum returns the number of handlers
func (l *Logger) HandlersNum() int { _ = "STUB: not implemented"; return 0 }

// LastErr get, will clear it after read.
func (l *Logger) LastErr() error { _ = "STUB: not implemented"; return nil }

//
// ---------------------------------------------------------------------------
// region Register handlers, processors
// ---------------------------------------------------------------------------
//

// AddHandler to the logger
func (l *Logger) AddHandler(h Handler) {
	_ = "STUB: not implemented"

	// AddHandlers to the logger
	return
}

func (l *Logger) AddHandlers(hs ...Handler) { _ = "STUB: not implemented"; return }

// PushHandler to the l. alias of AddHandler()
func (l *Logger) PushHandler(h Handler) {
	_ = "STUB: not implemented"

	// PushHandlers to the logger
	return
}

func (l *Logger) PushHandlers(hs ...Handler) { _ = "STUB: not implemented"; return }

// SetHandlers for the logger
func (l *Logger) SetHandlers(hs []Handler) {
	_ = "STUB: not implemented"

	// AddProcessor to the logger
	return
}

func (l *Logger) AddProcessor(p Processor) { _ = "STUB: not implemented"; return }

// PushProcessor to the logger, alias of AddProcessor()
func (l *Logger) PushProcessor(p Processor) { _ = "STUB: not implemented"; return }

// AddProcessors to the logger. alias of AddProcessor()
func (l *Logger) AddProcessors(ps ...Processor) { _ = "STUB: not implemented"; return }

// SetProcessors for the logger
func (l *Logger) SetProcessors(ps []Processor) {
	_ = "STUB: not implemented"

	// -------------------------- New sub-logger -----------------------------
	return
}

// NewSub return a new sub logger on the logger, can keep fields/data/ctx for sub logger.
//
// Usage:
//
//	sl := logger.NewSub().KeepCtx(custom ctx).
//		KeepFields(slog.M{"ip": ...}).
//		KeepData(slog.M{"username": ...})
//	defer sl.Release()
//
//	sl.Info("some message")
//	sl.Warn("some message")
func (l *Logger) NewSub() *SubLogger { _ = "STUB: not implemented"; return nil }

//
// ---------------------------------------------------------------------------
// region New record with logger
// ---------------------------------------------------------------------------
//

// Record return a new record with logger, will release after writing log.
func (l *Logger) Record() *Record { _ = "STUB: not implemented"; return nil }

// Reused return a new record with logger, but it can be reused.
// if you want to release the record, please call the Record.Release() after write log.
//
// Usage:
//
//	r := logger.Reused()
//	defer r.Release()
//
//	// can write log multiple times
//	r.Info("some message1")
//	r.Warn("some message1")
func (l *Logger) Reused() *Record { _ = "STUB: not implemented"; return nil }

// WithField new record with field
//
// TIP: add field need config Formatter template fields.
func (l *Logger) WithField(name string, value any) *Record {
	_ = "STUB: not implemented"

	// defer l.releaseRecord(r)
	return nil
}

// WithFields new record with fields
//
// TIP: add field need config Formatter template fields.
func (l *Logger) WithFields(fields M) *Record {
	_ = "STUB: not implemented"

	// defer l.releaseRecord(r)
	return nil
}

// WithData new record with data
func (l *Logger) WithData(data M) *Record { _ = "STUB: not implemented"; return nil }

// WithValue new record with data value
func (l *Logger) WithValue(key string, value any) *Record { _ = "STUB: not implemented"; return nil }

// WithExtra new record with extra data
func (l *Logger) WithExtra(ext M) *Record { _ = "STUB: not implemented"; return nil }

// WithTime new record with time.Time
func (l *Logger) WithTime(t time.Time) *Record {
	_ = "STUB: not implemented"

	// defer l.releaseRecord(r)
	return nil
}

// WithCtx new record with context.Context
func (l *Logger) WithCtx(ctx context.Context) *Record { _ = "STUB: not implemented"; return nil }

// WithContext new record with context.Context
func (l *Logger) WithContext(ctx context.Context) *Record {
	_ = "STUB: not implemented"

	// defer l.releaseRecord(r)
	return nil
}

//
// ---------------------------------------------------------------------------
// region Add log message
// ---------------------------------------------------------------------------
//

func (l *Logger) log(level Level, args []any) { _ = "STUB: not implemented"; return }

// Logf a format message with level
func (l *Logger) logf(level Level, format string, args []any) { _ = "STUB: not implemented"; return }

// logCtx a context message with level
func (l *Logger) logCtx(ctx context.Context, level Level, args []any) {
	_ = "STUB: not implemented"
	return
}

// logfCtx a format message with level,  context
func (l *Logger) logfCtx(ctx context.Context, level Level, format string, args []any) {
	_ = "STUB: not implemented"
	return
}

// Log a message with level
func (l *Logger) Log(level Level, args ...any) {
	_ = "STUB: not implemented"

	// Logf a format message with level
	return
}

func (l *Logger) Logf(level Level, format string, args ...any) { _ = "STUB: not implemented"; return }

// Print logs a message at level PrintLevel
func (l *Logger) Print(args ...any) { _ = "STUB: not implemented"; return }

// Println logs a message at level PrintLevel
func (l *Logger) Println(args ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message at level PrintLevel
func (l *Logger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Trace logs a message at level trace
func (l *Logger) Trace(args ...any) { _ = "STUB: not implemented"; return }

// Tracef logs a message at level trace
func (l *Logger) Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

// TraceCtx logs a message at level trace with context
func (l *Logger) TraceCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// TracefCtx logs a message at level trace with context
func (l *Logger) TracefCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Debug logs a message at level debug
func (l *Logger) Debug(args ...any) { _ = "STUB: not implemented"; return }

// Debugf logs a message at level debug
func (l *Logger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

// DebugCtx logs a message at level debug with context
func (l *Logger) DebugCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// DebugfCtx logs a message at level debug with context
func (l *Logger) DebugfCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Info logs a message at level Info
func (l *Logger) Info(args ...any) { _ = "STUB: not implemented"; return }

// Infof logs a message at level Info
func (l *Logger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

// InfoCtx logs a message at level Info with context
func (l *Logger) InfoCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// InfofCtx logs a message at level Info with context
func (l *Logger) InfofCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Notice logs a message at level notice
func (l *Logger) Notice(args ...any) { _ = "STUB: not implemented"; return }

// Noticef logs a message at level notice
func (l *Logger) Noticef(format string, args ...any) { _ = "STUB: not implemented"; return }

// NoticeCtx logs a message at level notice with context
func (l *Logger) NoticeCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// NoticefCtx logs a message at level notice with context
func (l *Logger) NoticefCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Warn logs a message at level Warn
func (l *Logger) Warn(args ...any) { _ = "STUB: not implemented"; return }

// Warnf logs a message at level Warn
func (l *Logger) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

// WarnCtx logs a message at level Warn with context
func (l *Logger) WarnCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// WarnfCtx logs a message at level Warn with context
func (l *Logger) WarnfCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Warning logs a message at level Warn, alias of Logger.Warn()
func (l *Logger) Warning(args ...any) { _ = "STUB: not implemented"; return }

// Error logs a message at level error
func (l *Logger) Error(args ...any) { _ = "STUB: not implemented"; return }

// Errorf logs a message at level error
func (l *Logger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

// ErrorT logs an error type at level error
func (l *Logger) ErrorT(err error) { _ = "STUB: not implemented"; return }

// ErrorCtx logs a message at level error with context
func (l *Logger) ErrorCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// ErrorfCtx logs a message at level error with context
func (l *Logger) ErrorfCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Stack logs a error message and with call stack. TODO
// func EStack(args ...any) { std.log(ErrorLevel, args) }

// Fatal logs a message at level fatal
func (l *Logger) Fatal(args ...any) { _ = "STUB: not implemented"; return }

// Fatalf logs a message at level fatal
func (l *Logger) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Fatalln logs a message at level fatal
func (l *Logger) Fatalln(args ...any) { _ = "STUB: not implemented"; return }

// FatalCtx logs a message at level panic with context
func (l *Logger) FatalCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// FatalfCtx logs a message at level panic with context
func (l *Logger) FatalfCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

// Panic logs a message at level panic
func (l *Logger) Panic(args ...any) { _ = "STUB: not implemented"; return }

// Panicf logs a message at level panic
func (l *Logger) Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Panicln logs a message at level panic
func (l *Logger) Panicln(args ...any) { _ = "STUB: not implemented"; return }

// PanicCtx logs a message at level panic with context
func (l *Logger) PanicCtx(ctx context.Context, args ...any) { _ = "STUB: not implemented"; return }

// PanicfCtx logs a message at level panic with context
func (l *Logger) PanicfCtx(ctx context.Context, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}
