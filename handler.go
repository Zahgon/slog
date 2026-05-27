package slog

import (
	"io"
)

//
// Handler interface
//

// Handler interface definition
type Handler interface {
	// Closer Close handler.
	// You should first call Flush() on close logic.
	// Refer the FileHandler.Close() handle
	io.Closer
	// Flush and sync logs to disk file.
	Flush() error
	// IsHandling Checks whether the given record will be handled by this handler.
	IsHandling(level Level) bool
	// Handle a log record.
	//
	// All records may be passed to this method, and the handler should discard
	// those that it does not want to handle.
	Handle(*Record) error
}

// LevelFormattable support limit log levels and provide formatter
type LevelFormattable interface {
	Formattable
	IsHandling(level Level) bool
}

// FormattableHandler interface
type FormattableHandler interface {
	Handler
	Formattable
}

/********************************************************************************
 * Common parts for handler
 ********************************************************************************/

// LevelWithFormatter struct definition
//
// - support set log formatter
// - only support set max log level
type LevelWithFormatter struct {
	FormattableTrait
	// Level max for logging messages. if current level <= Level will log messages
	Level Level
}

// NewLvFormatter create new LevelWithFormatter instance
func NewLvFormatter(maxLv Level) *LevelWithFormatter { _ = "STUB: not implemented"; return nil }

// SetMaxLevel set max level for logging messages
func (h *LevelWithFormatter) SetMaxLevel(maxLv Level) {
	_ = "STUB: not implemented"

	// IsHandling Check if the current level can be handling
	return
}

func (h *LevelWithFormatter) IsHandling(level Level) bool { _ = "STUB: not implemented"; return false }

// LevelsWithFormatter struct definition
//
// - support set log formatter
// - support setting multi log levels
type LevelsWithFormatter struct {
	FormattableTrait
	// Levels for logging messages
	Levels []Level
}

// NewLvsFormatter create new instance
func NewLvsFormatter(levels []Level) *LevelsWithFormatter { _ = "STUB: not implemented"; return nil }

// SetLimitLevels set limit levels for log message
func (h *LevelsWithFormatter) SetLimitLevels(levels []Level) {
	_ = "STUB: not implemented"

	// IsHandling Check if the current level can be handling
	return
}

func (h *LevelsWithFormatter) IsHandling(level Level) bool { _ = "STUB: not implemented"; return false }

// LevelMode define level mode for logging
type LevelMode uint8

// MarshalJSON implement the JSON Marshal interface [encoding/json.Marshaler]
func (m LevelMode) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implement the JSON Unmarshal interface [encoding/json.Unmarshaler]
func (m *LevelMode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// String return string value
func (m LevelMode) String() string { _ = "STUB: not implemented"; return "" }

const (
	// LevelModeList use level list for limit record write
	LevelModeList LevelMode = iota
	// LevelModeMax use max level limit log record write
	LevelModeMax
)

// SafeToLevelMode parse string value to LevelMode, fail return LevelModeList
func SafeToLevelMode(s string) LevelMode { _ = "STUB: not implemented"; return *new(LevelMode) }

// StringToLevelMode parse string value to LevelMode
func StringToLevelMode(s string) (LevelMode, error) {
	_ = "STUB: not implemented"
	return *new(LevelMode), nil
}

// is int value, try to parse as int

// LevelHandling struct definition
type LevelHandling struct {
	// level check mode. default is LevelModeList
	lvMode LevelMode
	// max level for a log message. if the current level <= Level will log a message
	maxLevel Level
	// levels limit for log message
	levels []Level
}

// SetMaxLevel set max level for a log message
func (h *LevelHandling) SetMaxLevel(maxLv Level) { _ = "STUB: not implemented"; return }

// SetLimitLevels set limit levels for log message
func (h *LevelHandling) SetLimitLevels(levels []Level) { _ = "STUB: not implemented"; return }

// IsHandling Check if the current level can be handling
func (h *LevelHandling) IsHandling(level Level) bool { _ = "STUB: not implemented"; return false }

// LevelFormatting wrap level handling and log formatter
type LevelFormatting struct {
	LevelHandling
	FormatterWrapper
}

// NewMaxLevelFormatting create new instance with max level
func NewMaxLevelFormatting(maxLevel Level) *LevelFormatting { _ = "STUB: not implemented"; return nil }

// NewLevelsFormatting create new instance with levels
func NewLevelsFormatting(levels []Level) *LevelFormatting { _ = "STUB: not implemented"; return nil }
