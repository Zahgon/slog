package slog

import "runtime"

//
// Formatter interface
//

// Formatter interface
type Formatter interface {
	// Format you can format record and write result to record.Buffer
	Format(record *Record) ([]byte, error)
}

// FormatterFunc wrapper definition
type FormatterFunc func(r *Record) ([]byte, error)

// Format a log record
func (fn FormatterFunc) Format(r *Record) ([]byte, error) {
	_ = "STUB: not implemented"

	// Formattable interface
	return nil, nil
}

type Formattable interface {
	// Formatter get the log formatter
	Formatter() Formatter
	// SetFormatter set the log formatter
	SetFormatter(Formatter)
}

// FormattableTrait alias of FormatterWrapper
type FormattableTrait = FormatterWrapper

// FormatterWrapper use for format log record.
//
// Default will use the TextFormatter
type FormatterWrapper struct {
	// if not set, default uses the TextFormatter
	formatter Formatter
}

// Formatter get formatter. if not set, will return TextFormatter
func (f *FormatterWrapper) Formatter() Formatter { _ = "STUB: not implemented"; return *new(Formatter) }

// SetFormatter to handler
func (f *FormatterWrapper) SetFormatter(formatter Formatter) { _ = "STUB: not implemented"; return }

// Format log record to bytes
func (f *FormatterWrapper) Format(record *Record) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CallerFormatFn caller format func
type CallerFormatFn func(rf *runtime.Frame) (cs string)

// AsTextFormatter util func
func AsTextFormatter(f Formatter) *TextFormatter { _ = "STUB: not implemented"; return nil }

// AsJSONFormatter util func
func AsJSONFormatter(f Formatter) *JSONFormatter { _ = "STUB: not implemented"; return nil }
