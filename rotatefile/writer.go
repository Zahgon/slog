package rotatefile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gookit/goutil/fsutil"
)

// Writer a flush, close, writer and support rotate file.
//
// refer https://github.com/flike/golog/blob/master/filehandler.go
type Writer struct {
	// writer instance id, use for debug
	id string
	mu sync.RWMutex
	// config of the writer
	cfg *Config

	// current opened logfile
	file *os.File
	// current opened file path. NOTE it maybe not equals Config.Filepath
	path string
	// The original file dir path for the Config.Filepath
	fileDir string
	// The original name and ext information
	fileName, onlyName, fileExt string

	// logfile max backup time. equals Config.BackupTime * time.Hour
	backupDur time.Duration
	// oldFiles []string
	cleanCh chan struct{}
	stopCh  chan struct{}

	// context use for rotating file by size
	written   uint64 // written size
	rotateNum uint   // rotate times number

	// ---- context use for rotating file by time ----

	// the rotating file name suffix format. eg: "20210102", "20210102_1500"
	suffixFormat   string
	checkInterval  int64     // check interval seconds.
	nextRotatingAt time.Time // next rotating time
}

// NewWriter create rotate write with config and init it.
func NewWriter(c *Config) (*Writer, error) { _ = "STUB: not implemented"; return nil, nil }

// NewWriterWith create a rotated writer with some settings.
func NewWriterWith(fns ...ConfigFn) (*Writer, error) { _ = "STUB: not implemented"; return nil, nil }

// init rotate dispatcher
func (d *Writer) init() error {
	d.id = fmt.Sprintf("%p", d)

	logfile := d.cfg.Filepath
	// dirSep := filepath.Separator
	// d.fileDir = filepath.Dir(logfile)
	d.fileDir, d.fileName = filepath.Split(d.cfg.Filepath)
	d.fileExt = filepath.Ext(d.fileName)                   // eg: .log
	d.onlyName = strings.TrimSuffix(d.fileName, d.fileExt) // eg: error
	// removes the trailing separator on the dir path
	if ln := len(d.fileDir); ln > 1 && d.fileDir[ln-1] == filepath.Separator {
		d.fileDir = d.fileDir[:ln-1]
	}

	d.backupDur = d.cfg.backupDuration()
	d.suffixFormat = d.cfg.RotateTime.TimeFormat()
	d.checkInterval = d.cfg.RotateTime.Interval()

	// calc and storage next rotating time
	if d.checkInterval > 0 {
		now := d.cfg.TimeClock.Now()
		// next rotating time
		d.nextRotatingAt = d.cfg.RotateTime.FirstCheckTime(now)
		if d.cfg.RotateMode == ModeCreate {
			// logfile = d.cfg.Filepath + "." + now.Format(d.suffixFormat)
			logfile = d.buildFilePath(now.Format(d.suffixFormat))
		}
	}

	// open the current file
	return d.openFile(logfile)
}

// Config gets the config
func (d *Writer) Config() Config {
	_ = "STUB: not implemented"

	// Flush sync data to disk. alias of Sync()
	return *new(Config)
}

func (d *Writer) Flush() error { _ = "STUB: not implemented"; return nil }

// Sync data to disk.
func (d *Writer) Sync() error { _ = "STUB: not implemented"; return nil }

// Close the writer. will sync data to disk, then close the file handle.
// and it will stop the async clean backups.
func (d *Writer) Close() error { _ = "STUB: not implemented"; return nil }

// MustClose the writer. alias of Close(), but will panic if has error.
func (d *Writer) MustClose() { _ = "STUB: not implemented"; return }

func (d *Writer) close(closeStopCh bool) error { _ = "STUB: not implemented"; return nil }

// stop the async clean backups

//
// ---------------------------------------------------------------------------
// write and rotate file
// ---------------------------------------------------------------------------
//

// WriteString implements the io.StringWriter
func (d *Writer) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0,

		// Write data to file. then check and do rotate file, then async clean backups
		nil
}

func (d *Writer) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// do write data
	return 0, nil
}

// do rotate file

// async clean backup files.

func (d *Writer) doWrite(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// if enable lock
	return 0, nil
}

// update size

// Rotate the file by config and async clean backups
func (d *Writer) Rotate() error {
	_ = "STUB: not implemented"

	// async clean backup files.
	return nil
}

// do rotate the logfile by config
func (d *Writer) doRotate() (err error) {
	_ = "STUB: not implemented"
	// if enable lock
	return nil
}

// do rotate a file by size

// do rotate a file by time

// TIP: should only call on d.checkInterval > 0
func (d *Writer) rotatingByTime() error { _ = "STUB: not implemented"; return nil }

// generate new file path.
// eg: /tmp/error.log => /tmp/error.20220423_1600.log
// file := d.cfg.Filepath + "." + d.nextRotatingAt.Format(d.suffixFormat)

// calc and storage next rotating time

func (d *Writer) rotatingBySize() error { _ = "STUB: not implemented"; return nil }

// up: use now minutes + seconds as rotate number

// eg: /tmp/error.log => /tmp/error.894136.log
// eg: /tmp/error.20220423_1600.log => /tmp/error.20220423_1600_894136.log

// rename current to new file by custom RenameFunc
// eg: /tmp/error.log => /tmp/error.163021_894136.log

// eg: /tmp/error.log => /tmp/error.25031615_894136.log

// always rename current to a new file

// rotateFile closes the syncBuffer's file and starts a new one.
func (d *Writer) rotatingFile(bakFile string, rename bool) error {
	_ = "STUB: not implemented"
	// close the current file
	return nil
}

// record old files for clean.
// d.oldFiles = append(d.oldFiles, bakFile)

// rename current to a new file.

// filepath for reopening

// reopen log file

// reset written

//
// ---------------------------------------------------------------------------
// clean backup files
// ---------------------------------------------------------------------------
//

// check should clean old files by config
func (d *Writer) shouldClean(withRand bool) bool { _ = "STUB: not implemented"; return false }

// 20% probability trigger clean

// async clean old files by config. should be in lock.
func (d *Writer) asyncClean() { _ = "STUB: not implemented"; return }

// if already running, send a signal

// add lock for deny concurrent clean

// re-check d.cleanCh is not nil

// init clean channel

// start a goroutine to clean backups

// consume the signal until stop

// stop clean

func (d *Writer) notifyClean() { _ = "STUB: not implemented"; return }

// notify clean old files

// skip on blocking

// Clean old files by config
func (d *Writer) Clean() (err error) { _ = "STUB: not implemented"; return nil }

// up: 单独运行清理，不需要设置 skipSeconds

// do clean old files by config
//
// - skipSeconds: skip find files that are within the specified seconds
func (d *Writer) doClean(skipSeconds ...int) (err error) {
	_ = "STUB: not implemented"
	// oldFiles: xx.log.yy files, no gz file
	return nil
}

// FIX: do not process recent changes to avoid conflicts

// find and clean old files

// fix: exclude the current file

// remove old gz files

// remove old log files

// remove old gz files
func (d *Writer) removeOldGzFiles(remNum int, gzFiles []fileInfo) (rn int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// sort by mod-time

// remove old log files
func (d *Writer) removeOldFiles(remNum int, oldFiles []fileInfo) (files []fileInfo, err error) {
	_ = "STUB: not implemented"
	// sort by mod-time, oldest at first.
	return nil, nil
}

//
// ---------------------------------------------------------------------------
// helper methods
// ---------------------------------------------------------------------------
//

// open the current file. and set the d.file, d.path
func (d *Writer) openFile(logfile string) error { _ = "STUB: not implemented"; return nil }

// return eg. logs/error.20220423_1600.log
func (d *Writer) buildFilePath(suffix string) string { _ = "STUB: not implemented"; return "" }

func (d *Writer) buildFilterFns(fileName string) []fsutil.FilterFunc {
	_ = "STUB: not implemented"
	return nil
}

// filter by name. match pattern like: error.log.* eg: error.log.xx, error.log.xx.gz

// ok, _ := path.Match(fileName+".*", ent.Name())

// 自定义文件名 eg: error.log -> error.20220423_02.log

// filter by mod-time, clear expired files

// skip, not handle

// collect un-expired

// remove expired files

func (d *Writer) compressFiles(oldFiles []fileInfo) error { _ = "STUB: not implemented"; return nil }

// remove an old log file

// Debug print debug message on development
func (d *Writer) debugLog(vs ...any) { _ = "STUB: not implemented"; return }
