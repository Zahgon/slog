package slog

import (
	"runtime"

	"github.com/valyala/bytebufferpool"
)

// const (
// 	defaultMaxCallerDepth  int = 15
// 	defaultKnownSlogFrames int = 4
// )

// Stack that attempts to recover the data for all goroutines.
// func getCallStacks(callerSkip int) []byte {
// 	return nil
// }

// FormatLevelName Format the level name, specify the length returned,
// fill the space with less length, and truncate than the length
func FormatLevelName(name string, length int) string { _ = "STUB: not implemented"; return "" }

func buildLowerLevelName() map[Level]string { _ = "STUB: not implemented"; return nil }

// getCaller retrieves the name of the first non-slog calling function
func getCaller(callerSkip int) (fr runtime.Frame, ok bool) {
	_ = "STUB: not implemented"
	return *
	// alloc 1 times
	new(runtime.Frame), false
}

func formatCaller(rf *runtime.Frame, flag uint8, userFn CallerFormatFn) (cs string) {
	_ = "STUB: not implemented"
	return ""
}

// CallerFlagFpLine

var msgBufPool bytebufferpool.Pool

// it like Println, will add spaces for each argument
func formatArgsWithSpaces(vs []any) string { _ = "STUB: not implemented"; return "" }

// cast is string, return it. NOT ALLOC MEMORY

// buf = make([]byte, 0, ln*8)

// TIP:
// `float` to string - will alloc 2 times memory
// `int <0`, `int > 100` to string -  will alloc 1 times memory

// add space

// return byteutil.String(bb.B) // perf: Reduce one memory allocation

// EncodeToString data to string
func EncodeToString(v any) string { _ = "STUB: not implemented"; return "" }

func mapToString(mp map[string]any) string { _ = "STUB: not implemented"; return "" }

// TODO use bytebufferpool

// remove last ', '

func parseTemplateToFields(tplStr string) []string { _ = "STUB: not implemented"; return nil }

func printStderr(args ...any) { _ = "STUB: not implemented"; return }
