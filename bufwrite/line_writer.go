package bufwrite

import (
	"io"
)

const (
	defaultBufSize = 1024 * 8
)

// LineWriter implements buffering for an io.Writer object.
// If an error occurs writing to a LineWriter, no more data will be
// accepted and all subsequent writes, and Flush, will return the error.
// After all data has been written, the client should call the
// Flush method to guarantee all data has been forwarded to
// the underlying io.Writer.
//
// from bufio.Writer.
//
// Change:
//
// always keep write full line. more difference please see Write
type LineWriter struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

// NewLineWriterSize returns a new LineWriter whose buffer has at least the specified
// size. If the argument io.Writer is already a LineWriter with large enough
// size, it returns the underlying LineWriter.
func NewLineWriterSize(w io.Writer, size int) *LineWriter {
	_ = "STUB: not implemented"
	// Is it already a LineWriter?
	return nil
}

// NewLineWriter returns a new LineWriter whose buffer has the default size.
func NewLineWriter(w io.Writer) *LineWriter { _ = "STUB: not implemented"; return nil }

// Size returns the size of the underlying buffer in bytes.
func (b *LineWriter) Size() int {
	_ = "STUB: not implemented"

	// Reset discards any unflushed buffered data, clears any error, and
	// resets b to write its output to w.
	return 0
}

func (b *LineWriter) Reset(w io.Writer) { _ = "STUB: not implemented"; return }

// Close implements the io.Closer
func (b *LineWriter) Close() error { _ = "STUB: not implemented"; return nil }

// is closer

// Sync implements the Syncer
func (b *LineWriter) Sync() error {
	_ = "STUB: not implemented"

	// Flush writes any buffered data to the underlying io.Writer.
	//
	// TIP: please add lock before calling the method.
	return nil
}

func (b *LineWriter) Flush() error { _ = "STUB: not implemented"; return nil }

// Available returns how many bytes are unused in the buffer.
func (b *LineWriter) Available() int { _ = "STUB: not implemented"; return 0 }

// Buffered returns the number of bytes that have been written into the current buffer.
func (b *LineWriter) Buffered() int {
	_ = "STUB: not implemented"

	// Write writes the contents of p into the buffer.
	// It returns the number of bytes written.
	// If nn < len(p), it also returns an error explaining
	// why the writing is short.
	return 0
}

func (b *LineWriter) Write(p []byte) (nn int, err error) {
	_ = "STUB: not implemented"
	// NOTE: 原来的 bufio.Writer#Write 会造成 p 写了一部分到 b.wr, 还有一部分在 b.buf，
	// 如果现在外部工具从 b.wr 收集数据，会收集到一行无法解析的数据(例如每个p是一行json日志)
	//
	//	for len(p) > b.Available() && b.err == nil {
	//		var n int
	//		if b.Buffered() == 0 {
	//			// Large write, empty buffer.
	//			// Write directly from p to avoid copy.
	//			n, b.err = b.wr.Write(p)
	//		} else {
	//			n = copy(b.buf[b.n:], p)
	//			b.n += n
	//			b.Flush()
	//		}
	//		nn += n
	//		p = p[n:]
	//	}
	return 0, nil
}

// UP: 改造一下逻辑，如果 len(p) > b.Available() 就将buf 和 p 都写入 b.wr

// WriteString to the writer
func (b *LineWriter) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }
