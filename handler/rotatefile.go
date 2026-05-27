package handler

import (
	"github.com/gookit/slog/rotatefile"
)

// NewRotateFileHandler instance. It supports splitting log files by time and size
func NewRotateFileHandler(logfile string, rt rotatefile.RotateTime, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustRotateFile handler instance, will panic on create error
func MustRotateFile(logfile string, rt rotatefile.RotateTime, fns ...ConfigFn) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewRotateFile instance. alias of NewRotateFileHandler()
func NewRotateFile(logfile string, rt rotatefile.RotateTime, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// ---------------------------------------------------------------------------
// rotate file by size
// ---------------------------------------------------------------------------
//

// MustSizeRotateFile instance
func MustSizeRotateFile(logfile string, maxSize int, fns ...ConfigFn) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewSizeRotateFile instance
func NewSizeRotateFile(logfile string, maxSize int, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSizeRotateFileHandler instance, default close rotate by time.
func NewSizeRotateFileHandler(logfile string, maxSize int, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	// close rotate by time.
	return nil, nil
}

//
// ---------------------------------------------------------------------------
// rotate log file by time
// ---------------------------------------------------------------------------
//

// RotateTime rotate log file by time.
//
// EveryDay:
//   - "error.log.20201223"
//
// EveryHour, Every30Minutes, EveryMinute:
//   - "error.log.20201223_1500"
//   - "error.log.20201223_1530"
//   - "error.log.20201223_1523"
//
// Deprecated: please use rotatefile.RotateTime
type RotateTime = rotatefile.RotateTime

// Deprecated: Please use define constants on pkg rotatefile. e.g. rotatefile.EveryDay
const (
	EveryDay  = rotatefile.EveryDay
	EveryHour = rotatefile.EveryDay

	Every30Minutes = rotatefile.Every30Min
	Every15Minutes = rotatefile.Every15Min

	EveryMinute = rotatefile.EveryMinute
	EverySecond = rotatefile.EverySecond // only use for tests
)

// MustTimeRotateFile instance
func MustTimeRotateFile(logfile string, rt rotatefile.RotateTime, fns ...ConfigFn) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewTimeRotateFile instance
func NewTimeRotateFile(logfile string, rt rotatefile.RotateTime, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTimeRotateFileHandler instance, default close rotate by size
func NewTimeRotateFileHandler(logfile string, rt rotatefile.RotateTime, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	// default close rotate by size: WithMaxSize(0)
	return nil, nil
}
