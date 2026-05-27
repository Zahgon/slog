package rotatefile

import (
	"time"
)

const defaultCheckInterval = 60 * time.Second

// CConfig struct for clean files
type CConfig struct {
	// BackupNum max number for keep old files.
	//
	// 0 is not limit, default is 20.
	BackupNum uint `json:"backup_num" yaml:"backup_num"`

	// BackupTime max time for keep old files, unit is TimeUnit.
	//
	// 0 is not limit, default is a week.
	BackupTime uint `json:"backup_time" yaml:"backup_time"`

	// Compress determines if the rotated log files should be compressed using gzip.
	// The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`

	// Patterns dir path with filename match patterns.
	//
	// eg: ["/tmp/error.log.*", "/path/to/info.log.*", "/path/to/dir/*"]
	Patterns []string `json:"patterns" yaml:"patterns"`

	// TimeClock for clean files
	TimeClock Clocker

	// TimeUnit for BackupTime. default is hours: time.Hour
	TimeUnit time.Duration `json:"time_unit" yaml:"time_unit"`

	// CheckInterval for clean files on daemon run. default is 60s.
	CheckInterval time.Duration `json:"check_interval" yaml:"check_interval"`

	// IgnoreError ignore remove error
	// TODO IgnoreError bool

	// RotateMode for rotate split files TODO
	//  - copy+cut: copy contents then truncate file
	//	- rename : rename file(use for like PHP-FPM app)
	// RotateMode RotateMode `json:"rotate_mode" yaml:"rotate_mode"`
}

// CConfigFunc for clean config
type CConfigFunc func(c *CConfig)

// AddDirPath for clean, will auto append * for match all files
func (c *CConfig) AddDirPath(dirPaths ...string) *CConfig { _ = "STUB: not implemented"; return nil }

// AddPattern for clean. eg: "/tmp/error.log.*"
func (c *CConfig) AddPattern(patterns ...string) *CConfig { _ = "STUB: not implemented"; return nil }

// WithConfigFn for custom settings
func (c *CConfig) WithConfigFn(fns ...CConfigFunc) *CConfig { _ = "STUB: not implemented"; return nil }

// NewCConfig instance
func NewCConfig() *CConfig { _ = "STUB: not implemented"; return nil }

// check interval time

// FilesClear multi files by time.
//
// use for rotate and clear other program produce log files
type FilesClear struct {
	// mu sync.Mutex
	cfg *CConfig
	// inited mark
	inited bool

	// file max backup time. equals CConfig.BackupTime * CConfig.TimeUnit
	backupDur  time.Duration
	quitDaemon chan struct{}
}

// NewFilesClear instance
func NewFilesClear(fns ...CConfigFunc) *FilesClear { _ = "STUB: not implemented"; return nil }

// Config get
func (r *FilesClear) Config() *CConfig {
	_ = "STUB: not implemented"

	// WithConfig for custom set config
	return nil
}

func (r *FilesClear) WithConfig(cfg *CConfig) *FilesClear { _ = "STUB: not implemented"; return nil }

// WithConfigFn for custom settings
func (r *FilesClear) WithConfigFn(fns ...CConfigFunc) *FilesClear {
	_ = "STUB: not implemented"
	return nil
}

//
// ---------------------------------------------------------------------------
// clean backup files
// ---------------------------------------------------------------------------
//

// StopDaemon for stop daemon clean
func (r *FilesClear) StopDaemon() { _ = "STUB: not implemented"; return }

// DaemonClean daemon clean old files by config
//
// NOTE: this method will block current goroutine
//
// Usage:
//
//	fc := rotatefile.NewFilesClear(nil)
//	fc.WithConfigFn(func(c *rotatefile.CConfig) {
//		c.AddDirPath("./testdata")
//	})
//
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//
//	// start daemon
//	go fc.DaemonClean(func() {
//		wg.Done()
//	})
//
//	// wait for stop
//	wg.Wait()
func (r *FilesClear) DaemonClean(onStop func()) { _ = "STUB: not implemented"; return }

// do cleaning

// Clean old files by config
func (r *FilesClear) prepare() { _ = "STUB: not implemented"; return }

// check backup time

// Clean old files by config
func (r *FilesClear) Clean() error { _ = "STUB: not implemented"; return nil }

// clear by time, can also clean by number

// CleanByPattern clean files by pattern
func (r *FilesClear) cleanByPattern(filePattern string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// find and clean expired files

// not handle subdir TODO: support subdir

// collect not expired

// remove expired file

// clear by backup number.

// sort by mod-time, oldest at first.

func (r *FilesClear) remove(filePath string) (err error) { _ = "STUB: not implemented"; return nil }
