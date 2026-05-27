// Package handler provide useful common log handlers.
//
// eg: file, console, multi_file, rotate_file, stream, syslog, email
package handler

import (
	"io"
	"os"
	"sync"

	"github.com/gookit/slog"
)

// DefaultBufferSize sizes the buffer associated with each log file. It's large
// so that log records can accumulate without the logging thread blocking
// on disk I/O. The flushDaemon will block instead.
var DefaultBufferSize = 8 * 1024

var (
	// DefaultFilePerm perm and flags for create log file
	DefaultFilePerm os.FileMode = 0664
	// DefaultFileFlags for create/open file
	DefaultFileFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND
)

// FlushWriter is the interface satisfied by logging destinations.
type FlushWriter interface {
	Flush() error
	// Writer the output writer
	io.Writer
}

// FlushCloseWriter is the interface satisfied by logging destinations.
type FlushCloseWriter interface {
	Flush() error
	// WriteCloser the output writer
	io.WriteCloser
}

// SyncCloseWriter is the interface satisfied by logging destinations.
// such as os.File
type SyncCloseWriter interface {
	Sync() error
	// WriteCloser the output writer
	io.WriteCloser
}

/********************************************************************************
 * Common parts for handler
 ********************************************************************************/

// LevelWithFormatter struct definition
//
// - support set log formatter
// - only support set one log level
//
// Deprecated: please use slog.LevelWithFormatter instead.
type LevelWithFormatter = slog.LevelWithFormatter

// LevelsWithFormatter struct definition
//
// - support set log formatter
// - support setting multi log levels
//
// Deprecated: please use slog.LevelsWithFormatter instead.
type LevelsWithFormatter = slog.LevelsWithFormatter

// NopFlushClose no operation.
//
// provide empty Flush(), Close() methods, useful for tests.
type NopFlushClose struct{}

// Flush logs to disk
func (h *NopFlushClose) Flush() error {
	_ = "STUB: not implemented"

	// Close handler
	return nil
}

func (h *NopFlushClose) Close() error {
	_ = "STUB: not implemented"

	// LockWrapper struct
	return nil
}

type LockWrapper struct {
	sync.Mutex
	disable bool
}

// Lock it
func (lw *LockWrapper) Lock() { _ = "STUB: not implemented"; return }

// Unlock it
func (lw *LockWrapper) Unlock() { _ = "STUB: not implemented"; return }

// EnableLock enable lock
func (lw *LockWrapper) EnableLock(enable bool) { _ = "STUB: not implemented"; return }

// LockEnabled status
func (lw *LockWrapper) LockEnabled() bool {
	_ = "STUB: not implemented"

	// QuickOpenFile like os.OpenFile
	return false
}

func QuickOpenFile(filepath string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }
