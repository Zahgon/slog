package handler

import (
	"io"

	"github.com/gookit/slog"
)

// SyncCloseHandler definition
type SyncCloseHandler struct {
	slog.LevelFormattable
	Output SyncCloseWriter
}

// NewSyncCloserWithLF create new SyncCloseHandler, with custom slog.LevelFormattable
func NewSyncCloserWithLF(out SyncCloseWriter, lf slog.LevelFormattable) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// init formatter and level handle

//
// ------------- Use max log level -------------
//

// SyncCloserWithMaxLevel create new SyncCloseHandler, with max log level
func SyncCloserWithMaxLevel(out SyncCloseWriter, maxLevel slog.Level) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

//
// ------------- Use multi log levels -------------
//

// NewSyncCloser create new SyncCloseHandler, alias of NewSyncCloseHandler()
func NewSyncCloser(out SyncCloseWriter, levels []slog.Level) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// SyncCloserWithLevels create new SyncCloseHandler, alias of NewSyncCloseHandler()
func SyncCloserWithLevels(out SyncCloseWriter, levels []slog.Level) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewSyncCloseHandler create new SyncCloseHandler with limited log levels
//
// Usage:
//
//	f, err := os.OpenFile("my.log", ...)
//	h := handler.NewSyncCloseHandler(f, slog.AllLevels)
func NewSyncCloseHandler(out SyncCloseWriter, levels []slog.Level) *SyncCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// Close the handler
func (h *SyncCloseHandler) Close() error { _ = "STUB: not implemented"; return nil }

// Flush the handler
func (h *SyncCloseHandler) Flush() error { _ = "STUB: not implemented"; return nil }

// Writer of the handler
func (h *SyncCloseHandler) Writer() io.Writer {
	_ = "STUB: not implemented"

	// Handle log record
	return *new(io.Writer)
}

func (h *SyncCloseHandler) Handle(record *slog.Record) error { _ = "STUB: not implemented"; return nil }
