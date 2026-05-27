package main

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

// run: go run ./_example/slog_all_level.go
func main() {
	l := slog.NewWithConfig(func(l *slog.Logger) {
		l.DoNothingOnPanicFatal()
	})

	l.AddHandler(handler.NewConsoleHandler(slog.AllLevels))
	printAllLevel(l, "this is a", "log", "message")
}

func printAllLevel(l *slog.Logger, args ...any) { _ = "STUB: not implemented"; return }
