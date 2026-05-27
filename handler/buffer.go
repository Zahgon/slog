package handler

import (
	"io"
	"os"

	"github.com/gookit/slog"
)

// NewBuffered create new BufferedHandler
func NewBuffered(w io.WriteCloser, bufSize int, levels ...slog.Level) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewBufferedHandler create new BufferedHandler
func NewBufferedHandler(w io.WriteCloser, bufSize int, levels ...slog.Level) *FlushCloseHandler {
	_ = "STUB: not implemented"
	return nil
}

// LineBufferedFile handler
func LineBufferedFile(logfile string, bufSize int, levels []slog.Level) (slog.Handler, error) {
	_ = "STUB: not implemented"
	return *new(slog.Handler), nil
}

// LineBuffOsFile handler
func LineBuffOsFile(f *os.File, bufSize int, levels []slog.Level) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// LineBuffWriter handler
func LineBuffWriter(w io.Writer, bufSize int, levels []slog.Level) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

//
// --------- wrap a handler with buffer ---------
//

// FormatWriterHandler interface
type FormatWriterHandler interface {
	slog.Handler
	// Formatter record formatter
	Formatter() slog.Formatter
	// Writer the output writer
	Writer() io.Writer
}
