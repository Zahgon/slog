package handler

import (
	"io"

	"github.com/gookit/slog"
)

// WriteCloserHandler definition
type WriteCloserHandler struct {
	slog.LevelFormattable
	Output io.WriteCloser
}

// NewWriteCloserWithLF create new WriteCloserHandler and with custom slog.LevelFormattable
func NewWriteCloserWithLF(out io.WriteCloser, lf slog.LevelFormattable) *WriteCloserHandler {
	_ = "STUB: not implemented"
	return nil
}

// init formatter and level handle

// WriteCloserWithMaxLevel create new WriteCloserHandler and with max log level
func WriteCloserWithMaxLevel(out io.WriteCloser, maxLevel slog.Level) *WriteCloserHandler {
	_ = "STUB: not implemented"
	return nil
}

//
// ------------- Use multi log levels -------------
//

// WriteCloserWithLevels create a new instance and with limited log levels
func WriteCloserWithLevels(out io.WriteCloser, levels []slog.Level) *WriteCloserHandler {
	_ = "STUB: not implemented"
	// h := &WriteCloserHandler{Output: out}
	// h.LimitLevels(levels)
	return nil
}

// NewWriteCloser create a new instance
func NewWriteCloser(out io.WriteCloser, levels []slog.Level) *WriteCloserHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewWriteCloserHandler create new WriteCloserHandler
//
// Usage:
//
//	buf := new(bytes.Buffer)
//	h := handler.NewIOWriteCloserHandler(&buf, slog.AllLevels)
//
//	f, err := os.OpenFile("my.log", ...)
//	h := handler.NewIOWriteCloserHandler(f, slog.AllLevels)
func NewWriteCloserHandler(out io.WriteCloser, levels []slog.Level) *WriteCloserHandler {
	_ = "STUB: not implemented"
	return nil
}

// Close the handler
func (h *WriteCloserHandler) Close() error { _ = "STUB: not implemented"; return nil }

// Flush the handler
func (h *WriteCloserHandler) Flush() error {
	_ = "STUB: not implemented"

	// Handle log record
	return nil
}

func (h *WriteCloserHandler) Handle(record *slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}
