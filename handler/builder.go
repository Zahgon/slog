package handler

import (
	"io"

	"github.com/gookit/slog"
	"github.com/gookit/slog/rotatefile"
)

//
// ---------------------------------------------------------------------------
// handler builder
// ---------------------------------------------------------------------------
//

// Builder struct for create handler
type Builder struct {
	*Config
	Output io.Writer
}

// NewBuilder create
func NewBuilder() *Builder { _ = "STUB: not implemented"; return nil }

// WithOutput to the builder
func (b *Builder) WithOutput(w io.Writer) *Builder { _ = "STUB: not implemented"; return nil }

// With some config fn
//
// Deprecated: please use WithConfigFn()
func (b *Builder) With(fns ...ConfigFn) *Builder { _ = "STUB: not implemented"; return nil }

// WithConfigFn some config fn
func (b *Builder) WithConfigFn(fns ...ConfigFn) *Builder { _ = "STUB: not implemented"; return nil }

// WithLogfile setting
func (b *Builder) WithLogfile(logfile string) *Builder { _ = "STUB: not implemented"; return nil }

// WithLevelMode setting
func (b *Builder) WithLevelMode(mode slog.LevelMode) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithLogLevel setting max log level
func (b *Builder) WithLogLevel(level slog.Level) *Builder { _ = "STUB: not implemented"; return nil }

// WithLogLevels setting
func (b *Builder) WithLogLevels(levels []slog.Level) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithBuffMode setting
func (b *Builder) WithBuffMode(bufMode string) *Builder { _ = "STUB: not implemented"; return nil }

// WithBuffSize setting
func (b *Builder) WithBuffSize(bufSize int) *Builder { _ = "STUB: not implemented"; return nil }

// WithMaxSize setting
func (b *Builder) WithMaxSize(maxSize uint64) *Builder { _ = "STUB: not implemented"; return nil }

// WithRotateTime setting
func (b *Builder) WithRotateTime(rt rotatefile.RotateTime) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithCompress setting
func (b *Builder) WithCompress(compress bool) *Builder { _ = "STUB: not implemented"; return nil }

// WithUseJSON setting
func (b *Builder) WithUseJSON(useJSON bool) *Builder { _ = "STUB: not implemented"; return nil }

// Build slog handler.
func (b *Builder) Build() slog.FormattableHandler {
	_ = "STUB: not implemented"
	return *new(slog.FormattableHandler)
}

// Build slog handler.
func (b *Builder) buildFromWriter(w io.Writer) (h slog.FormattableHandler) {
	_ = "STUB: not implemented"
	return *new(slog.FormattableHandler)
}

// use json format.

// rest builder.
func (b *Builder) reset() { _ = "STUB: not implemented"; return }
