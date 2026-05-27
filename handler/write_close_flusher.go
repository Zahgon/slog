package handler

import (
	"github.com/gookit/slog"
)

// FlushCloseHandler definition
type FlushCloseHandler struct {
	slog.LevelFormattable
	Output FlushCloseWriter
}

// NewFlushCloserWithLF create new FlushCloseHandler, with custom slog.LevelFormattable
func NewFlushCloserWithLF(out FlushCloseWriter, lf slog.LevelFormattable) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// init formatter and level handle

//
// ------------- Use max log level -------------
//

// FlushCloserWithMaxLevel create new FlushCloseHandler, with max log level
func FlushCloserWithMaxLevel(out FlushCloseWriter, maxLevel slog.Level) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

//
// ------------- Use multi log levels -------------
//

// NewFlushCloser create new FlushCloseHandler, alias of NewFlushCloseHandler()
func NewFlushCloser(out FlushCloseWriter, levels []slog.Level) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// FlushCloserWithLevels create new FlushCloseHandler, alias of NewFlushCloseHandler()
func FlushCloserWithLevels(out FlushCloseWriter, levels []slog.Level) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewFlushCloseHandler create new FlushCloseHandler
//
// Usage:
//
//	buf := new(byteutil.Buffer)
//	h := handler.NewFlushCloseHandler(&buf, slog.AllLevels)
//
//	f, err := os.OpenFile("my.log", ...)
//	h := handler.NewFlushCloseHandler(f, slog.AllLevels)
func NewFlushCloseHandler(out FlushCloseWriter, levels []slog.Level) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// Close the handler
func (h *FlushCloseHandler) Close() error { _ = "STUB: not implemented"; return nil }

// Flush the handler
func (h *FlushCloseHandler) Flush() error { _ = "STUB: not implemented"; return nil }

// Handle log record
func (h *FlushCloseHandler) Handle(record *slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}
