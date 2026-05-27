package slog

import (
	"context"
	"runtime"
	"time"
)

// Record a log record definition
type Record struct {
	logger *Logger
	// reuse flag. for reuse a Record, will not be released on after writing.
	// release the record need call Release() method.
	//
	// Reuse field: Ctx, Data, Fields, Extra
	//
	// NOTE, if you reuse a record, you must call Reused() method.
	reuse bool
	// Mark whether the current record is released to the pool. TODO
	freed bool
	// inited flag for record
	inited bool

	// Time for record log, if is empty will use now.
	//
	// TIP: Will be emptied after each use (write)
	Time time.Time
	// Level log level for record
	Level Level
	// level name cache from Level
	levelName string

	// Channel log channel name. eg: "order", "goods", "user"
	Channel string
	Message string

	// Ctx context.Context
	Ctx context.Context

	// Fields custom fields data.
	// Contains all the fields set by the user.
	Fields M
	// Data log context data
	Data M
	// Extra log extra data
	Extra M

	// Caller information
	Caller *runtime.Frame
	// CallerFlag value. default is equals to Logger.CallerFlag
	CallerFlag uint8
	// CallerSkip value. default is equals to Logger.CallerSkip
	CallerSkip int
	// EnableStack enable stack info, default is false. TODO
	EnableStack bool

	// Buffer Can use Buffer on formatter
	// Buffer *bytes.Buffer

	// log input args backups, from log() and logf(). its dont use in formatter.
	Fmt  string
	Args []any
}

func newRecord(logger *Logger) *Record { _ = "STUB: not implemented"; return nil }

// with some options

// init map data field
// Data:   make(M, 2),
// Extra:  make(M, 0),
// Fields: make(M, 0),

// Reused set record is reused, will not be released on after writing.
// so, MUST call Release() method after use completed.
func (r *Record) Reused() *Record { _ = "STUB: not implemented"; return nil }

// Release manual release record to pool
func (r *Record) Release() { _ = "STUB: not implemented"; return }

//
// ---------------------------------------------------------------------------
// Copy record with something
// ---------------------------------------------------------------------------
//

// WithTime set the record time
func (r *Record) WithTime(t time.Time) *Record { _ = "STUB: not implemented"; return nil }

// WithCtx on record
func (r *Record) WithCtx(ctx context.Context) *Record { _ = "STUB: not implemented"; return nil }

// WithContext on record
func (r *Record) WithContext(ctx context.Context) *Record { _ = "STUB: not implemented"; return nil }

// WithError on record
func (r *Record) WithError(err error) *Record { _ = "STUB: not implemented"; return nil }

// WithData on record
func (r *Record) WithData(data M) *Record { _ = "STUB: not implemented"; return nil }

// WithField with a new field to record
//
// Note: add field need config Formatter template fields.
func (r *Record) WithField(name string, val any) *Record { _ = "STUB: not implemented"; return nil }

// WithFields with new fields to record
//
// Note: add field need config Formatter template fields.
func (r *Record) WithFields(fields M) *Record { _ = "STUB: not implemented"; return nil }

// Copy new record from old record
func (r *Record) Copy() *Record { _ = "STUB: not implemented"; return nil }

// reuse: true, // copy record is reused

// Time:       r.Time,

//
// ---------------------------------------------------------------------------
// Direct set value to record
// ---------------------------------------------------------------------------
//

// SetCtx on record
func (r *Record) SetCtx(ctx context.Context) *Record { _ = "STUB: not implemented"; return nil }

// SetContext on record
func (r *Record) SetContext(ctx context.Context) *Record { _ = "STUB: not implemented"; return nil }

// SetData on record
func (r *Record) SetData(data M) *Record { _ = "STUB: not implemented"; return nil }

// AddData on record
func (r *Record) AddData(data M) *Record { _ = "STUB: not implemented"; return nil }

// WithValue add Data value to record. alias of AddValue
func (r *Record) WithValue(key string, value any) *Record { _ = "STUB: not implemented"; return nil }

// AddValue add Data value to record
func (r *Record) AddValue(key string, value any) *Record { _ = "STUB: not implemented"; return nil }

// Value get Data value from record
func (r *Record) Value(key string) any { _ = "STUB: not implemented"; return *new(any) }

// SetExtra information on record
func (r *Record) SetExtra(data M) *Record { _ = "STUB: not implemented"; return nil }

// AddExtra information on record
func (r *Record) AddExtra(data M) *Record { _ = "STUB: not implemented"; return nil }

// SetExtraValue on record
func (r *Record) SetExtraValue(k string, v any) { _ = "STUB: not implemented"; return }

// SetTime on record
func (r *Record) SetTime(t time.Time) *Record { _ = "STUB: not implemented"; return nil }

// AddField add new field to the record
func (r *Record) AddField(name string, val any) *Record { _ = "STUB: not implemented"; return nil }

// AddFields add new fields to the record
func (r *Record) AddFields(fields M) *Record { _ = "STUB: not implemented"; return nil }

// SetFields to the record
func (r *Record) SetFields(fields M) *Record { _ = "STUB: not implemented"; return nil }

// Field value gets from record
func (r *Record) Field(key string) any { _ = "STUB: not implemented"; return *new(any) }

//
// ---------------------------------------------------------------------------
// Add log message with builder
// TODO r.Build(InfoLevel).Str().Int().Float().Msg()
// ---------------------------------------------------------------------------
//

// Object data on record TODO optimize performance
// func (r *Record) Obj(obj fmt.Stringer) *Record {
// 	r.Data = ctx
// 	return r
// }

// Object data on record TODO optimize performance
// func (r *Record) Any(v any) *Record {
// 	r.Data = ctx
// 	return r
// }

// func (r *Record) Str(message string) {
// 	r.logWrite(level, []byte(message))
// }

// func (r *Record) Int(val int) {
// 	r.logWrite(level, []byte(message))
// }

//
// ---------------------------------------------------------------------------
// Add log message with level
// ---------------------------------------------------------------------------
//

func (r *Record) log(level Level, args []any) { _ = "STUB: not implemented"; return }

// r.Message = strutil.Byte2str(formatArgsWithSpaces(args)) // will reduce memory allocation once

// do write log, then release record

func (r *Record) logf(level Level, format string, args []any) { _ = "STUB: not implemented"; return }

// do write log, then release record

// Log a message with level
func (r *Record) Log(level Level, args ...any) {
	_ = "STUB: not implemented"

	// Logf a message with level
	return
}

func (r *Record) Logf(level Level, format string, args ...any) { _ = "STUB: not implemented"; return }

// Info logs a message at level Info
func (r *Record) Info(args ...any) { _ = "STUB: not implemented"; return }

// Infof logs a message at level Info
func (r *Record) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

// Trace logs a message at level Trace
func (r *Record) Trace(args ...any) { _ = "STUB: not implemented"; return }

// Tracef logs a message at level Trace
func (r *Record) Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

// Error logs a message at level Error
func (r *Record) Error(args ...any) { _ = "STUB: not implemented"; return }

// Errorf logs a message at level Error
func (r *Record) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Warn logs a message at level Warn
func (r *Record) Warn(args ...any) { _ = "STUB: not implemented"; return }

// Warnf logs a message at level Warn
func (r *Record) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Notice logs a message at level Notice
func (r *Record) Notice(args ...any) { _ = "STUB: not implemented"; return }

// Noticef logs a message at level Notice
func (r *Record) Noticef(format string, args ...any) { _ = "STUB: not implemented"; return }

// Debug logs a message at level Debug
func (r *Record) Debug(args ...any) { _ = "STUB: not implemented"; return }

// Debugf logs a message at level Debug
func (r *Record) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Print logs a message at level Print
func (r *Record) Print(args ...any) { _ = "STUB: not implemented"; return }

// Println logs a message at level Print. alias of Print
func (r *Record) Println(args ...any) { _ = "STUB: not implemented"; return }

// Printf logs a message at level Print
func (r *Record) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Fatal logs a message at level Fatal
func (r *Record) Fatal(args ...any) { _ = "STUB: not implemented"; return }

// Fatalln logs a message at level Fatal
func (r *Record) Fatalln(args ...any) { _ = "STUB: not implemented"; return }

// Fatalf logs a message at level Fatal
func (r *Record) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// Panic logs a message at level Panic
func (r *Record) Panic(args ...any) { _ = "STUB: not implemented"; return }

// Panicln logs a message at level Panic
func (r *Record) Panicln(args ...any) { _ = "STUB: not implemented"; return }

// Panicf logs a message at level Panic
func (r *Record) Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }

// ---------------------------------------------------------------------------
// helper methods
// ---------------------------------------------------------------------------

// LevelName get
func (r *Record) LevelName() string {
	_ = "STUB: not implemented"

	// GoString of the record
	return ""
}

func (r *Record) GoString() string { _ = "STUB: not implemented"; return "" }

func (r *Record) timestamp() string { _ = "STUB: not implemented"; return "" }
