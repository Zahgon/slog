package rotatefile

import (
	"io/fs"
	"time"
)

const compressSuffix = ".gz"

func printErrln(pfx string, err error) { _ = "STUB: not implemented"; return }

func compressFile(srcPath, dstPath string) error { _ = "STUB: not implemented"; return nil }

// create and open a gz file

// do copy

// TODO replace to fsutil.FileInfo
type fileInfo struct {
	fs.FileInfo
	filePath string
}

// Path get file full path. eg: "/path/to/file.go"
func (fi *fileInfo) Path() string { _ = "STUB: not implemented"; return "" }

func newFileInfo(filePath string, fi fs.FileInfo) fileInfo {
	_ = "STUB: not implemented"
	return *new(fileInfo)
}

// modTimeFInfos sorts by oldest time modified in the fileInfo.
// eg: [old_220211, old_220212, old_220213]
type modTimeFInfos []fileInfo

// Less check
func (fis modTimeFInfos) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap value
func (fis modTimeFInfos) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Len get
func (fis modTimeFInfos) Len() int {
	_ = "STUB: not implemented"

	// MockClocker mock clock for test
	return 0
}

type MockClocker struct {
	tt time.Time
}

// NewMockClock create a mock time instance from datetime string.
func NewMockClock(datetime string) *MockClocker { _ = "STUB: not implemented"; return nil }

// Now get current time.
func (mt *MockClocker) Now() time.Time {
	_ = "STUB: not implemented"

	// Add progresses time by the given duration.
	return *new(time.Time)
}

func (mt *MockClocker) Add(d time.Duration) { _ = "STUB: not implemented"; return }

// Datetime returns the current time in the format "2006-01-02 15:04:05".
func (mt *MockClocker) Datetime() string { _ = "STUB: not implemented"; return "" }
