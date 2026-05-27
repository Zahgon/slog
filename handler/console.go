package handler

import (
	"github.com/gookit/slog"
)

/********************************************************************************
 * console log handler
 ********************************************************************************/

// ConsoleHandler definition
type ConsoleHandler = IOWriterHandler

// NewConsoleWithLF create new ConsoleHandler and with custom slog.LevelFormattable
func NewConsoleWithLF(lf slog.LevelFormattable) *ConsoleHandler {
	_ = "STUB: not implemented"
	return nil
}

// default use text formatter

// default enable color on console

//
// ------------- Use max log level -------------
//

// ConsoleWithMaxLevel create new ConsoleHandler and with max log level
func ConsoleWithMaxLevel(level slog.Level) *ConsoleHandler { _ = "STUB: not implemented"; return nil }

//
// ------------- Use multi log levels -------------
//

// NewConsole create new ConsoleHandler, alias of NewConsoleHandler
func NewConsole(levels []slog.Level) *ConsoleHandler { _ = "STUB: not implemented"; return nil }

// ConsoleWithLevels create new ConsoleHandler and with limited log levels
func ConsoleWithLevels(levels []slog.Level) *ConsoleHandler { _ = "STUB: not implemented"; return nil }

// NewConsoleHandler create new ConsoleHandler with limited log levels
func NewConsoleHandler(levels []slog.Level) *ConsoleHandler { _ = "STUB: not implemented"; return nil }
