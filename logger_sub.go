package slog

import "context"

// SubLogger is a sub-logger, It can be used to keep a certain amount of contextual information and log multiple times.
// 可以用于保持一定的上下文信息多次记录日志。例如在循环中使用，或者作为方法参数传入。
//
// Usage:
//
//	sl := slog.NewSub().KeepCtx(custom ctx).
//		KeepFields(slog.M{"ip": ...}).
//		KeepData(slog.M{"username": ...})
//	defer sl.Release()
//
//	sl.Info("some message")
type SubLogger struct {
	l *Logger // parent logger

	// Ctx keep context for all log records
	Ctx context.Context
	// Fields keep custom fields data for all log records
	Fields M
	// Data keep data for all log records
	Data M
	// Extra data. will keep for all log records
	Extra M
}

// NewSubWith returns a new SubLogger with parent logger.
func NewSubWith(l *Logger) *SubLogger { _ = "STUB: not implemented"; return nil }

// KeepCtx keep context for all log records
func (sub *SubLogger) KeepCtx(ctx context.Context) *SubLogger {
	_ = "STUB: not implemented"
	return nil

	// KeepFields keep custom fields data for all log records
}

func (sub *SubLogger) KeepFields(fields M) *SubLogger { _ = "STUB: not implemented"; return nil }

// KeepField keep custom field for all log records
func (sub *SubLogger) KeepField(field string, value any) *SubLogger {
	_ = "STUB: not implemented"
	return nil
}

// KeepData keep data for all log records
func (sub *SubLogger) KeepData(data M) *SubLogger { _ = "STUB: not implemented"; return nil }

// KeepExtra keep extra data for all log records
func (sub *SubLogger) KeepExtra(extra M) *SubLogger { _ = "STUB: not implemented"; return nil }

// Release releases the SubLogger.
func (sub *SubLogger) Release() { _ = "STUB: not implemented"; return }

func (sub *SubLogger) withKeepCtx() *Record { _ = "STUB: not implemented"; return nil }

//
// ---------------------------------------------------------------------------
// Add log message with level
// ---------------------------------------------------------------------------
//

// Print logs a message at PrintLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Print(args ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message at PrintLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Trace logs a message at TraceLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Trace(args ...any) { _ = "STUB: not implemented"; return }

// Tracef logs a formatted message at TraceLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

// Debug logs a message at DebugLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Debug(args ...any) { _ = "STUB: not implemented"; return }

// Debugf logs a formatted message at DebugLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Info logs a message at InfoLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Info(args ...any) { _ = "STUB: not implemented"; return }

// Infof logs a formatted message at InfoLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

// Notice logs a message at NoticeLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Notice(args ...any) { _ = "STUB: not implemented"; return }

// Noticef logs a formatted message at NoticeLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Noticef(format string, args ...any) { _ = "STUB: not implemented"; return }

// Warn logs a message at WarnLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Warn(args ...any) { _ = "STUB: not implemented"; return }

// Warnf logs a formatted message at WarnLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Error logs a message at ErrorLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Error(args ...any) { _ = "STUB: not implemented"; return }

// Errorf logs a formatted message at ErrorLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Fatal logs a message at FatalLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Fatal(args ...any) { _ = "STUB: not implemented"; return }

// Fatalf logs a formatted message at FatalLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Panic logs a message at PanicLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Panic(args ...any) { _ = "STUB: not implemented"; return }

// Panicf logs a formatted message at PanicLevel. will with sub logger's context, fields and data
func (sub *SubLogger) Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }
