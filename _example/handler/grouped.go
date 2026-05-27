package handler

import "github.com/gookit/slog"

/********************************************************************************
 * Grouped Handler
 ********************************************************************************/

// GroupedHandler definition
type GroupedHandler struct {
	handlers []slog.Handler
	// Levels for log message
	Levels []slog.Level
	// IgnoreErr on handling messages
	IgnoreErr bool
}

// NewGroupedHandler create new GroupedHandler
func NewGroupedHandler(handlers []slog.Handler) *GroupedHandler {
	_ = "STUB: not implemented"
	return nil
}

// IsHandling Check if the current level can be handling
func (h *GroupedHandler) IsHandling(level slog.Level) bool { _ = "STUB: not implemented"; return false }

// Handle log record
func (h *GroupedHandler) Handle(record *slog.Record) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Close log handlers
func (h *GroupedHandler) Close() error { _ = "STUB: not implemented"; return nil }

// Flush log records
func (h *GroupedHandler) Flush() error { _ = "STUB: not implemented"; return nil }
