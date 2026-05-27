package rotatefile

import (
	"os"
	"time"

	"github.com/gookit/goutil/timex"
)

//
// ---------------------------- rotate time -------------------------------
//

type rotateLevel uint8

const (
	levelDay rotateLevel = iota
	levelHour
	levelMin
	levelSec
)

// RotateTime for a rotating file. unit is seconds.
//
// EveryDay:
//   - "error.log.20201223"
//
// EveryHour, Every30Min, EveryMinute:
//   - "error.log.20201223_1500"
//   - "error.log.20201223_1530"
//   - "error.log.20201223_1523"
type RotateTime int

// built in rotate time constants
const (
	EveryMonth  RotateTime = 30 * timex.OneDaySec
	EveryDay    RotateTime = timex.OneDaySec
	EveryHour   RotateTime = timex.OneHourSec
	Every30Min  RotateTime = 30 * timex.OneMinSec
	Every15Min  RotateTime = 15 * timex.OneMinSec
	EveryMinute RotateTime = timex.OneMinSec
	EverySecond RotateTime = 1 // only use for tests
)

// Interval get check interval time. unit is seconds.
func (rt RotateTime) Interval() int64 {
	_ = "STUB: not implemented"

	// FirstCheckTime for a rotated file.
	// - will automatically align the time from the start of each hour.
	return 0
}

func (rt RotateTime) FirstCheckTime(now time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// should check on H:59:59.500

// eg: minutes=5

// will rotate at next hour start. eg: now.Minute()=57, nextMin=62.

// eg: now.Minute()=37, nextMin=42, will get nextDur=40

// levelSec

// level for rotating time
func (rt RotateTime) level() rotateLevel { _ = "STUB: not implemented"; return *new(rotateLevel) }

// TimeFormat get log file suffix format
//
// EveryDay:
//   - "error.log.20201223"
//
// EveryHour, Every30Min, EveryMinute:
//   - "error.log.20201223_1500"
//   - "error.log.20201223_1530"
//   - "error.log.20201223_1523"
func (rt RotateTime) TimeFormat() (suffixFormat string) { _ = "STUB: not implemented"; return "" }

// default is levelHour

// MarshalJSON implement the JSON Marshal interface [encoding/json.Marshaler]
func (rt RotateTime) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implement the JSON Unmarshal interface [encoding/json.Unmarshaler]
func (rt *RotateTime) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// String rotate type to string
func (rt RotateTime) String() string { _ = "STUB: not implemented"; return "" }

// levelSec

// StringToRotateTime parse and convert string to RotateTime
func StringToRotateTime(s string) (RotateTime, error) {
	_ = "STUB: not implemented"
	// is int value, try to parse as seconds
	return *new(RotateTime), nil
}

// parse time duration string. eg: "1h", "1m", "1d"

//
// ---------------------------- RotateMode -------------------------------
//

// RotateMode for a rotated file. 0: rename, 1: create
type RotateMode uint8

const (
	// ModeRename rotating file by rename.
	//
	// Example flow:
	//  - always write to "error.log"
	//  - rotating by rename it to "error.log.20201223"
	//  - then re-create "error.log"
	ModeRename RotateMode = iota

	// ModeCreate rotating file by create a new file.
	//
	// Example flow:
	//  - directly create a new file on each rotated time. eg: "error.log.20201223", "error.log.20201224"
	ModeCreate
)

// String get string name
func (m RotateMode) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implement the JSON Marshal interface [encoding/json.Marshaler]
func (m RotateMode) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implement the JSON Unmarshal interface [encoding/json.Unmarshaler]
func (m *RotateMode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// StringToRotateMode convert string to RotateMode
func StringToRotateMode(s string) (RotateMode, error) {
	_ = "STUB: not implemented"
	return *new(RotateMode), nil
}

// is int value, try to parse as int

//
// ---------------------------- Clocker -------------------------------
//

// Clocker is the interface used for determine the current time
type Clocker interface {
	Now() time.Time
}

// ClockFn func
type ClockFn func() time.Time

// Now implements the Clocker
func (fn ClockFn) Now() time.Time {
	_ = "STUB: not implemented"

	// ConfigFn for setting config
	return *new(time.Time)
}

type ConfigFn func(c *Config)

// Config struct for rotate dispatcher
type Config struct {
	// Filepath the log file path, will be rotating. eg: "logs/error.log"
	Filepath string `json:"filepath" yaml:"filepath"`

	// FilePerm for create log file. default DefaultFilePerm
	FilePerm os.FileMode `json:"file_perm" yaml:"file_perm"`

	// RotateMode for rotate file. default ModeRename
	RotateMode RotateMode `json:"rotate_mode" yaml:"rotate_mode"`

	// MaxSize file contents max size, unit is bytes.
	// If is equals zero, disable rotate file by size
	//
	// default see DefaultMaxSize
	MaxSize uint64 `json:"max_size" yaml:"max_size"`

	// RotateTime the file rotating interval time, unit is seconds.
	// If is equals zero, disable rotate file by time
	//
	// default: EveryHour
	RotateTime RotateTime `json:"rotate_time" yaml:"rotate_time"`

	// CloseLock use sync lock on writing contents, rotating file.
	//
	// default: false
	CloseLock bool `json:"close_lock" yaml:"close_lock"`

	// BackupNum max number for keep old files.
	//
	// 0 is not limit, default is DefaultBackNum
	BackupNum uint `json:"backup_num" yaml:"backup_num"`

	// BackupTime max time for keep old files, unit is hours.
	//
	// 0 is not limit, default is DefaultBackTime
	BackupTime uint `json:"backup_time" yaml:"backup_time"`

	// CleanOnClose determines if the rotated log files should be cleaned up when close.
	CleanOnClose bool `json:"clean_on_close" yaml:"clean_on_close"`

	// Compress determines if the rotated log files should be compressed using gzip.
	// The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`

	// RenameFunc you can custom-build filename for rotate file by size.
	//
	// Example:
	//
	//  c.RenameFunc = func(filepath string, rotateNum uint) string {
	// 		suffix := time.Now().Format("06010215")
	//
	// 		// eg: /tmp/error.log => /tmp/error.log.24032116_894136
	// 		return filepath + fmt.Sprintf(".%s_%d", suffix, rotateNum)
	//  }
	RenameFunc func(filePath string, rotateNum uint) string `json:"-" yaml:"-"`

	// TimeClock for a rotating file by time.
	TimeClock Clocker `json:"-" yaml:"-"`

	// DebugMode for debug on development.
	DebugMode bool `json:"debug_mode" yaml:"debug_mode"`
}

func (c *Config) backupDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// With more config setting func
func (c *Config) With(fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

// Create new Writer by config
func (c *Config) Create() (*Writer, error) {
	_ = "STUB: not implemented"

	// IsMode check rotate mode
	return nil, nil
}

func (c *Config) IsMode(m RotateMode) bool { _ = "STUB: not implemented"; return false }

var (
	// DefaultFilePerm perm and flags for create log file
	DefaultFilePerm os.FileMode = 0664
	// DefaultFileFlags for open log file
	DefaultFileFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND

	// DefaultTimeClockFn for create time
	DefaultTimeClockFn = ClockFn(func() time.Time {
		return time.Now()
	})
)

// NewDefaultConfig instance
func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// RenameFunc: DefaultFilenameFn,

// NewConfig by file path, and can with custom setting
func NewConfig(filePath string, fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWith custom func
func NewConfigWith(fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

// EmptyConfigWith new empty config with custom func
func EmptyConfigWith(fns ...ConfigFn) *Config {
	_ = "STUB: not implemented"

	// RenameFunc: DefaultFilenameFn,
	return nil
}

// WithFilepath setting
func WithFilepath(logfile string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithDebugMode setting for debug mode
func WithDebugMode(c *Config) {
	_ = "STUB: not implemented"

	// WithCompress setting for compress
	return
}

func WithCompress(c *Config) {
	_ = "STUB: not implemented"

	// WithBackupNum setting for backup number
	return
}

func WithBackupNum(num uint) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }
