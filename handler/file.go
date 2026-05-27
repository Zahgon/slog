package handler

import (
	"github.com/gookit/slog"
)

// JSONFileHandler create new FileHandler with JSON formatter
func JSONFileHandler(logfile string, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBuffFileHandler create file handler with buff size
func NewBuffFileHandler(logfile string, buffSize int, fns ...ConfigFn) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustFileHandler create file handler
func MustFileHandler(logfile string, fns ...ConfigFn) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewFileHandler create new FileHandler
func NewFileHandler(logfile string, fns ...ConfigFn) (h *SyncCloseHandler, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// ------------- simple file handler -------------
//

// MustSimpleFile new instance
func MustSimpleFile(filepath string, maxLv ...slog.Level) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewSimpleFile new instance
func NewSimpleFile(filepath string, maxLv ...slog.Level) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSimpleFileHandler instance, default log level is InfoLevel
//
// Usage:
//
//	h, err := NewSimpleFileHandler("/tmp/error.log")
//
// Custom formatter:
//
//	h.SetFormatter(slog.NewJSONFormatter())
//	slog.PushHandler(h)
//	slog.Info("log message")
func NewSimpleFileHandler(filePath string, maxLv ...slog.Level) (*SyncCloseHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
