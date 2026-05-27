package slog

import (
	"github.com/gookit/color"
	"github.com/valyala/bytebufferpool"
)

// there are built in text log template
const (
	DefaultTemplate = "[{{datetime}}] [{{channel}}] [{{level}}] [{{caller}}] {{message}} {{data}} {{extra}}\n"
	NamedTemplate   = "{{datetime}} channel={{channel}} level={{level}} [file={{caller}}] message={{message}} data={{data}}\n"
)

// ColorTheme for format log to console
var ColorTheme = map[Level]color.Color{
	PanicLevel:  color.FgRed,
	FatalLevel:  color.FgRed,
	ErrorLevel:  color.FgMagenta,
	WarnLevel:   color.FgYellow,
	NoticeLevel: color.OpBold,
	InfoLevel:   color.FgGreen,
	DebugLevel:  color.FgCyan,
	// TraceLevel:  color.FgLightGreen,
}

// TextFormatter definition
type TextFormatter struct {
	// template text template for render output log messages
	template string
	// fields list, parsed from template string.
	//
	// NOTE: contains no-field items in the list. eg: ["level", "}}"}
	fields []string

	// TimeFormat the time format layout. default is DefaultTimeFormat
	TimeFormat string
	// Enable color on print log to terminal
	EnableColor bool
	// ColorTheme setting on render color on terminal
	ColorTheme map[Level]color.Color
	// FullDisplay Whether to display when record.Data, record.Extra, etc. are empty
	FullDisplay bool
	// EncodeFunc data encode for Record.Data, Record.Extra, etc.
	//
	// Default is encode by EncodeToString()
	EncodeFunc func(v any) string
	// CallerFormatFunc the caller format layout. default is defined by CallerFlag
	CallerFormatFunc CallerFormatFn
	// LevelFormatFunc custom the level name format.
	LevelFormatFunc func(s string) string
	// ColorRenderFunc custom color render func.
	//
	// - `s`: level name or message
	ColorRenderFunc func(filed, s string, l Level) string

	// TODO BeforeFunc call it before format, update fields or other
	// BeforeFunc func(r *Record)
}

// TextFormatterFn definition
type TextFormatterFn func(*TextFormatter)

// NewTextFormatter create new TextFormatter
func NewTextFormatter(template ...string) *TextFormatter { _ = "STUB: not implemented"; return nil }

// default options

// EnableColor: color.SupportColor(),
// EncodeFunc: func(v any) string {
// 	return fmt.Sprint(v)
// },

// TextFormatterWith create new TextFormatter with options
func TextFormatterWith(fns ...TextFormatterFn) *TextFormatter {
	_ = "STUB: not implemented"
	return nil
}

// LimitLevelNameLen limit the length of the level name
func LimitLevelNameLen(length int) TextFormatterFn {
	_ = "STUB: not implemented"
	return *new(TextFormatterFn)
}

// Configure the formatter
func (f *TextFormatter) Configure(fn TextFormatterFn) *TextFormatter {
	_ = "STUB: not implemented"
	return nil

	// WithOptions func on the formatter
}

func (f *TextFormatter) WithOptions(fns ...TextFormatterFn) *TextFormatter {
	_ = "STUB: not implemented"
	return nil
}

// SetTemplate set the log format template and update field-map
func (f *TextFormatter) SetTemplate(fmtTpl string) { _ = "STUB: not implemented"; return }

// Template get
func (f *TextFormatter) Template() string {
	_ = "STUB: not implemented"

	// WithEnableColor enable color on print log to terminal
	return ""
}

func (f *TextFormatter) WithEnableColor(enable bool) *TextFormatter {
	_ = "STUB: not implemented"
	return nil
}

// Fields get an export field list
func (f *TextFormatter) Fields() []string { _ = "STUB: not implemented"; return nil }

var textPool bytebufferpool.Pool

// Format a log record
//
//goland:noinspection GoUnhandledErrorResult
func (f *TextFormatter) Format(r *Record) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// record formatted custom fields

// is not field name. eg: "}}] "

// remove left "}}"

// UP: check not configured fields in template.

// return buf.Bytes(), nil

func (f *TextFormatter) beforeFormat() {
	_ = "STUB: not implemented"
	// if f.BeforeFunc == nil {}
	return
}

func (f *TextFormatter) renderColorText(field, s string, l Level) string {
	_ = "STUB: not implemented"
	// custom level name format
	return ""
}

// custom color render func

// output colored logs for console output
