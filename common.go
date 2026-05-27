package slog

import (
	"time"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/gsr"
)

// SLogger interface
type SLogger interface {
	gsr.Logger
	Log(level Level, v ...any)
	Logf(level Level, format string, v ...any)
}

// LoggerFn func
type LoggerFn func(l *Logger)

//
// log level definitions
// region Log level

// Level type
type Level uint32

// String get level name
func (l Level) String() string {
	_ = "STUB: not implemented"

	// Name get level name. eg: INFO, DEBUG ...
	return ""
}

func (l Level) Name() string {
	_ = "STUB: not implemented"

	// LowerName get lower level name. eg: info, debug ...
	return ""
}

func (l Level) LowerName() string { _ = "STUB: not implemented"; return "" }

// ShouldHandling compare level, if current level <= l, it will be record.
func (l Level) ShouldHandling(curLevel Level) bool { _ = "STUB: not implemented"; return false }

// MarshalJSON implement the JSON Marshal interface [encoding/json.Marshaler]
func (l Level) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implement the JSON Unmarshal interface [encoding/json.Unmarshaler]
func (l *Level) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Levels level list
type Levels []Level

// Contains given level
func (ls Levels) Contains(level Level) bool { _ = "STUB: not implemented"; return false }

// These are the different logging levels. You can set the logging level to log handler
const (
	// PanicLevel level, the highest level of severity. will call panic() if the logging level <= PanicLevel.
	PanicLevel Level = 100
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level <= FatalLevel.
	FatalLevel Level = 200
	// ErrorLevel level. Runtime errors. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel Level = 300
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel Level = 400
	// NoticeLevel level Uncommon events
	NoticeLevel Level = 500
	// InfoLevel level. Examples: User logs in, SQL logs.
	InfoLevel Level = 600
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel Level = 700
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel Level = 800
)

//
// some common definitions
// region common types

// StringMap string map short name
type StringMap = map[string]string

// M short name of map[string]any
type M map[string]any

// String map to string
func (m M) String() string { _ = "STUB: not implemented"; return "" }

// ClockFn func
type ClockFn func() time.Time

// Now implements the Clocker
func (fn ClockFn) Now() time.Time {
	_ = "STUB: not implemented"

	// region CallerFlagMode
	return *new(time.Time)
}

// CallerFlagMode Defines the Caller backtrace information mode.
type CallerFlagMode = uint8

// NOTICE: you must set `Logger.ReportCaller=true` for reporting caller.
// then config the Logger.CallerFlag by follow flags.
const (
	// CallerFlagFnlFcn report short func name with filename and with line.
	// eg: "logger_test.go:48,TestLogger_ReportCaller"
	CallerFlagFnlFcn CallerFlagMode = iota
	// CallerFlagFull full func name with filename and with line.
	// eg: "github.com/gookit/slog_test.TestLogger_ReportCaller(),logger_test.go:48"
	CallerFlagFull
	// CallerFlagFunc full package with func name.
	// eg: "github.com/gookit/slog_test.TestLogger_ReportCaller"
	CallerFlagFunc
	// CallerFlagFcLine full package with func name and with line.
	// eg: "github.com/gookit/slog_test.TestLogger_ReportCaller:48"
	CallerFlagFcLine
	// CallerFlagPkg report full package name.
	// eg: "github.com/gookit/slog_test"
	CallerFlagPkg
	// CallerFlagPkgFnl report full package name + filename + line.
	// eg: "github.com/gookit/slog_test,logger_test.go:48"
	CallerFlagPkgFnl
	// CallerFlagFpLine report full filepath with line.
	// eg: "/work/go/gookit/slog/logger_test.go:48"
	CallerFlagFpLine
	// CallerFlagFnLine report filename with line.
	// eg: "logger_test.go:48"
	CallerFlagFnLine
	// CallerFlagFcName only report func name.
	// eg: "TestLogger_ReportCaller"
	CallerFlagFcName
)

var (
	// FieldKeyData define the key name for Record.Data
	FieldKeyData = "data"
	// FieldKeyTime key name
	FieldKeyTime = "time"
	// FieldKeyDate key name
	FieldKeyDate = "date"

	// FieldKeyDatetime key name
	FieldKeyDatetime = "datetime"
	// FieldKeyTimestamp key name
	FieldKeyTimestamp = "timestamp"

	// FieldKeyCaller the field key name for report caller.
	//
	// For caller style please see CallerFlagFull, CallerFlagFunc and more.
	//
	// NOTICE: you must set `Logger.ReportCaller=true` for reporting caller
	FieldKeyCaller = "caller"

	// FieldKeyLevel name
	FieldKeyLevel = "level"
	// FieldKeyError Define the key when adding errors using WithError.
	FieldKeyError = "error"
	// FieldKeyExtra key name
	FieldKeyExtra = "extra"

	// FieldKeyChannel name
	FieldKeyChannel = "channel"
	// FieldKeyMessage name
	FieldKeyMessage = "message"
)

// region Global variables

var (
	// DefaultChannelName for log record
	DefaultChannelName = "application"
	// DefaultTimeFormat define
	DefaultTimeFormat = "2006/01/02T15:04:05.000"

	// DebugMode enable debug mode for logger. use for local development.
	DebugMode = envutil.GetBool("OPEN_SLOG_DEBUG", false)

	// DoNothingOnExit handle func. use for testing.
	DoNothingOnExit = func(code int) {}
	// DoNothingOnPanic handle func. use for testing.
	DoNothingOnPanic = func(v any) {}

	// DefaultPanicFn handle func
	DefaultPanicFn = func(v any) { panic(v) }
	// DefaultClockFn create func
	DefaultClockFn = ClockFn(func() time.Time { return time.Now() })
)

var (
	// PrintLevel for use Logger.Print / Printf / Println
	PrintLevel = InfoLevel

	// AllLevels exposing all logging levels
	AllLevels = Levels{
		PanicLevel,
		FatalLevel,
		ErrorLevel,
		WarnLevel,
		NoticeLevel,
		InfoLevel,
		DebugLevel,
		TraceLevel,
	}

	// DangerLevels define the commonly danger log levels
	DangerLevels = Levels{PanicLevel, FatalLevel, ErrorLevel, WarnLevel}
	// NormalLevels define the commonly normal log levels
	NormalLevels = Levels{InfoLevel, NoticeLevel, DebugLevel, TraceLevel}

	// LevelNames all level mapping name
	LevelNames = map[Level]string{
		PanicLevel:  "PANIC",
		FatalLevel:  "FATAL",
		ErrorLevel:  "ERROR",
		WarnLevel:   "WARNING",
		NoticeLevel: "NOTICE",
		InfoLevel:   "INFO",
		DebugLevel:  "DEBUG",
		TraceLevel:  "TRACE",
	}

	// lower level name.
	lowerLevelNames = buildLowerLevelName()
	// empty time for reset record.
	emptyTime = time.Time{}
)

// region Global functions

// LevelName match
func LevelName(l Level) string { _ = "STUB: not implemented"; return "" }

// LevelByName convert name to level, fallback to InfoLevel if not match
func LevelByName(ln string) Level { _ = "STUB: not implemented"; return *new(Level) }

// Name2Level convert name to level
func Name2Level(s string) (Level, error) {
	_ = "STUB: not implemented"
	return *

	// StringToLevel parse and convert string value to Level
	new(Level), nil
}

func StringToLevel(s string) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }

// make the zero value useful

// is int value, try to parse as int

//
// exit handle logic
//

// global exit handler
var exitHandlers = make([]func(), 0)

func runExitHandlers() { _ = "STUB: not implemented"; return }

// ExitHandlers get all global exitHandlers
func ExitHandlers() []func() { _ = "STUB: not implemented"; return nil }

// RegisterExitHandler register an exit-handler on global exitHandlers
func RegisterExitHandler(handler func()) { _ = "STUB: not implemented"; return }

// PrependExitHandler prepend register an exit-handler on global exitHandlers
func PrependExitHandler(handler func()) { _ = "STUB: not implemented"; return }

// ResetExitHandlers reset all exitHandlers
func ResetExitHandlers(applyToStd bool) { _ = "STUB: not implemented"; return }
