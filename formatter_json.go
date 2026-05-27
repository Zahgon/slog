package slog

import (
	"github.com/valyala/bytebufferpool"
)

var (
	// DefaultFields default log export fields for json formatter.
	DefaultFields = []string{
		FieldKeyDatetime,
		FieldKeyChannel,
		FieldKeyLevel,
		FieldKeyCaller,
		FieldKeyMessage,
		FieldKeyData,
		FieldKeyExtra,
	}

	// NoTimeFields log export fields without time
	NoTimeFields = []string{
		FieldKeyChannel,
		FieldKeyLevel,
		FieldKeyMessage,
		FieldKeyData,
		FieldKeyExtra,
	}
)

// JSONFormatter definition
type JSONFormatter struct {
	// Fields set exported common log fields. default is DefaultFields
	Fields []string
	// Aliases for output fields. you can change the export field name.
	//
	// - item: `"field" : "output name"`
	//
	// eg: {"message": "msg"} export field will display "msg"
	Aliases StringMap

	// PrettyPrint will indent all JSON logs
	PrettyPrint bool
	// TimeFormat the time format layout. default is DefaultTimeFormat
	TimeFormat string
	// CallerFormatFunc the caller format layout. default is defined by CallerFlag
	CallerFormatFunc CallerFormatFn
}

// NewJSONFormatter create new JSONFormatter
func NewJSONFormatter(fn ...func(f *JSONFormatter)) *JSONFormatter {
	_ = "STUB: not implemented"
	return nil

	// Aliases: make(StringMap, 0),
}

// Configure current formatter
func (f *JSONFormatter) Configure(fn func(*JSONFormatter)) *JSONFormatter {
	_ = "STUB: not implemented"

	// AddField for export
	return nil
}

func (f *JSONFormatter) AddField(name string) *JSONFormatter { _ = "STUB: not implemented"; return nil }

var jsonPool bytebufferpool.Pool

// Format a log record to JSON bytes
func (f *JSONFormatter) Format(r *Record) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO perf: use buf write build JSON string.

// default:
// 	logData[outName] = r.Fields[field]

// exported custom record fields

// sort.Interface()

// buf.Reset()

// buf := r.NewBuffer()
// buf.Reset()
// buf.Grow(256)

// has been added newline in Encode().
