package handler

import (
	"io"
	"io/fs"

	"github.com/gookit/slog"
	"github.com/gookit/slog/rotatefile"
)

// the buff mode constants
const (
	BuffModeLine = "line"
	BuffModeBite = "bite"
)

const (
	// LevelModeList use level list for limit record write
	LevelModeList = slog.LevelModeList
	// LevelModeValue use max level limit log record write
	LevelModeValue = slog.LevelModeMax
)

// ConfigFn for config some settings
type ConfigFn func(c *Config)

// Config struct
type Config struct {
	// Logfile for writing logs
	Logfile string `json:"logfile" yaml:"logfile"`

	// FilePerm for create log file. default rotatefile.DefaultFilePerm
	FilePerm fs.FileMode `json:"file_perm" yaml:"file_perm"`

	// LevelMode for limit log records. default LevelModeList
	LevelMode slog.LevelMode `json:"level_mode" yaml:"level_mode"`

	// Level max value. valid on LevelMode = LevelModeValue
	//
	// eg: set Level=slog.LevelError, it will only write messages on level <= error.
	Level slog.Level `json:"level" yaml:"level"`

	// Levels list for writing. valid on LevelMode = LevelModeList
	Levels []slog.Level `json:"levels" yaml:"levels"`

	// UseJSON for format logs
	UseJSON bool `json:"use_json" yaml:"use_json"`

	// BuffMode type name. allow: line, bite
	//
	// Recommend use BuffModeLine(it's default)
	BuffMode string `json:"buff_mode" yaml:"buff_mode"`

	// BuffSize for enable buffer, unit is bytes. set 0 to disable buffer
	BuffSize int `json:"buff_size" yaml:"buff_size"`

	// RotateTime for a rotating file, unit is seconds.
	RotateTime rotatefile.RotateTime `json:"rotate_time" yaml:"rotate_time"`

	// RotateMode for a rotating file by time. default rotatefile.ModeRename
	RotateMode rotatefile.RotateMode `json:"rotate_mode" yaml:"rotate_mode"`

	// TimeClock for a rotating file by time.
	TimeClock rotatefile.Clocker `json:"-" yaml:"-"`

	// MaxSize for a rotating file by size, unit is bytes.
	MaxSize uint64 `json:"max_size" yaml:"max_size"`

	// Compress determines if the rotated log files should be compressed using gzip.
	// The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`

	// BackupNum max number for keep old files.
	//
	// 0 is not limit, default is 20.
	BackupNum uint `json:"backup_num" yaml:"backup_num"`

	// BackupTime max time for keep old files, unit is hours.
	//
	// 0 is not limit, default is a week.
	BackupTime uint `json:"backup_time" yaml:"backup_time"`

	// RenameFunc build filename for rotate file
	RenameFunc func(filepath string, rotateNum uint) string

	// CleanOnClose determines if the rotated log files should be cleaned up when close.
	CleanOnClose bool `json:"clean_on_close" yaml:"clean_on_close"`

	// DebugMode for debug on development.
	DebugMode bool
}

// NewEmptyConfig new config instance
func NewEmptyConfig(fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

// NewConfig new config instance with some default settings.
func NewConfig(fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

// rotate file settings

// old files clean settings

// FromJSON load config from json string
func (c *Config) FromJSON(bts []byte) error { _ = "STUB: not implemented"; return nil }

// With more config settings func
func (c *Config) With(fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

// WithConfigFn more config settings func
func (c *Config) WithConfigFn(fns ...ConfigFn) *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) newLevelFormattable() slog.LevelFormattable {
	_ = "STUB: not implemented"
	return *new(slog.LevelFormattable)
}

// CreateHandler quick create a handler by config
func (c *Config) CreateHandler() (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// with log level and formatter

// RotateWriter build rotate writer by config
func (c *Config) RotateWriter() (output SyncCloseWriter, err error) {
	_ = "STUB: not implemented"
	return *new(SyncCloseWriter), nil
}

// CreateWriter build writer by config
func (c *Config) CreateWriter() (output SyncCloseWriter, err error) {
	_ = "STUB: not implemented"
	return *new(SyncCloseWriter), nil
}

// create a rotated writer by config.

// has locked on logger.write()

// copy settings

// create a file writer

// wrap buffer

type flushSyncCloseWriter interface {
	FlushCloseWriter
	Sync() error
}

// wrap buffer for the writer
func (c *Config) wrapBuffer(w io.Writer) (bw flushSyncCloseWriter) {
	_ = "STUB: not implemented"
	return *new(flushSyncCloseWriter)
}

//
// ---------------------------------------------------------------------------
// global config func
// ---------------------------------------------------------------------------
//

// WithLogfile setting
func WithLogfile(logfile string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithFilePerm setting
func WithFilePerm(filePerm fs.FileMode) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLevelMode setting
func WithLevelMode(lm slog.LevelMode) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLevelModeString setting
func WithLevelModeString(s string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLogLevel setting max log level
func WithLogLevel(level slog.Level) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLevelName setting max level by name
func WithLevelName(name string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithMaxLevelName setting max level by name
func WithMaxLevelName(name string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLogLevels setting
func WithLogLevels(levels slog.Levels) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLevelNamesString setting multi levels by level names string, multi names split by comma.
func WithLevelNamesString(names string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithLevelNames set multi levels by level names.
func WithLevelNames(names []string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithRotateTime setting the rotated time
func WithRotateTime(rt rotatefile.RotateTime) ConfigFn {
	_ = "STUB: not implemented"
	return *new(ConfigFn)
}

// WithRotateTimeString setting the rotated time by string.
//
// eg: "1hour", "24h", "1day", "7d", "1m", "30s"
func WithRotateTimeString(s string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithRotateMode setting rotating mode rotatefile.RotateMode
func WithRotateMode(m rotatefile.RotateMode) ConfigFn {
	_ = "STUB: not implemented"
	return *new(ConfigFn)
}

// WithRotateModeString setting rotatefile.RotateMode by string.
func WithRotateModeString(s string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithTimeClock setting
func WithTimeClock(clock rotatefile.Clocker) ConfigFn {
	_ = "STUB: not implemented"
	return *new(ConfigFn)
}

// WithBackupNum setting
func WithBackupNum(n uint) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithBackupTime setting backup time
func WithBackupTime(bt uint) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithBuffMode setting buffer mode
func WithBuffMode(buffMode string) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithBuffSize setting buffer size, unit is bytes.
func WithBuffSize(buffSize int) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithMaxSize setting max size for a rotated file
func WithMaxSize(maxSize uint64) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithCompress setting compress
func WithCompress(compress bool) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithUseJSON setting uses JSON format
func WithUseJSON(useJSON bool) ConfigFn { _ = "STUB: not implemented"; return *new(ConfigFn) }

// WithDebugMode setting for debug mode
func WithDebugMode(c *Config) { _ = "STUB: not implemented"; return }
