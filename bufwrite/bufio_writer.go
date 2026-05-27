// Package bufwrite provides buffered io.Writer with sync and close methods.
package bufwrite

import (
	"bufio"
	"io"
)

// BufIOWriter wrap the bufio.Writer, implements the Sync() Close() methods
type BufIOWriter struct {
	bufio.Writer
	// backup the bufio.Writer.wr
	writer io.Writer
}

// NewBufIOWriterSize instance with size
func NewBufIOWriterSize(w io.Writer, size int) *BufIOWriter { _ = "STUB: not implemented"; return nil }

// NewBufIOWriter instance
func NewBufIOWriter(w io.Writer) *BufIOWriter { _ = "STUB: not implemented"; return nil }

// Close implements the io.Closer
func (w *BufIOWriter) Close() error { _ = "STUB: not implemented"; return nil }

// is closer

// Sync implements the Syncer
func (w *BufIOWriter) Sync() error { _ = "STUB: not implemented"; return nil }
