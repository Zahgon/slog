package handler

import (
	"io"

	"github.com/gookit/slog"
)

// IOWriterHandler definition
type IOWriterHandler struct {
	NopFlushClose
	slog.LevelFormattable
	Output io.Writer
}

// TextFormatter get the formatter
func (h *IOWriterHandler) TextFormatter() *slog.TextFormatter {
	_ = "STUB: not implemented"
	return nil
}

// Handle log record
func (h *IOWriterHandler) Handle(record *slog.Record) error { _ = "STUB: not implemented"; return nil }

// NewIOWriterWithLF create new IOWriterHandler, with custom slog.LevelFormattable
func NewIOWriterWithLF(out io.Writer, lf slog.LevelFormattable) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}

// init formatter and level handle

//
// ------------- Use max log level -------------
//

// IOWriterWithMaxLevel create new IOWriterHandler, with max log level
//
// Usage:
//
//		buf := new(bytes.Buffer)
//		h := handler.IOWriterWithMaxLevel(buf, slog.InfoLevel)
//	 slog.AddHandler(h)
//		slog.Info("info message")
func IOWriterWithMaxLevel(out io.Writer, maxLevel slog.Level) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}

//
// ------------- Use multi log levels -------------
//

// NewIOWriter create a new instance and with limited log levels
func NewIOWriter(out io.Writer, levels []slog.Level) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}

// IOWriterWithLevels create a new instance and with limited log levels
func IOWriterWithLevels(out io.Writer, levels []slog.Level) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewIOWriterHandler create new IOWriterHandler
//
// Usage:
//
//	buf := new(bytes.Buffer)
//	h := handler.NewIOWriterHandler(&buf, slog.AllLevels)
//
//	f, err := os.OpenFile("my.log", ...)
//	h := handler.NewIOWriterHandler(f, slog.AllLevels)
func NewIOWriterHandler(out io.Writer, levels []slog.Level) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}

// SimpleHandler definition. alias of IOWriterHandler
type SimpleHandler = IOWriterHandler

// NewHandler create a new instance
func NewHandler(out io.Writer, maxLevel slog.Level) *SimpleHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewSimple create a new instance
func NewSimple(out io.Writer, maxLevel slog.Level) *SimpleHandler {
	_ = "STUB: not implemented"
	return nil
}

// SimpleWithLevels create new simple handler, with log levels
func SimpleWithLevels(out io.Writer, levels []slog.Level) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewSimpleHandler create new SimpleHandler
//
// Usage:
//
//	buf := new(bytes.Buffer)
//	h := handler.NewSimpleHandler(&buf, slog.InfoLevel)
//
//	f, err := os.OpenFile("my.log", ...)
//	h := handler.NewSimpleHandler(f, slog.InfoLevel)
func NewSimpleHandler(out io.Writer, maxLevel slog.Level) *IOWriterHandler {
	_ = "STUB: not implemented"
	return nil
}
