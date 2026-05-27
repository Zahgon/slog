//go:build !windows && !plan9

package handler

import (
	"log/syslog"

	"github.com/gookit/slog"
)

// SysLogOpt for syslog handler
type SysLogOpt struct {
	// Tag syslog tag
	Tag string
	// Priority syslog priority
	Priority syslog.Priority
	// Network syslog network
	Network string
	// Raddr syslog address
	Raddr string
}

// SysLogHandler struct
type SysLogHandler struct {
	slog.LevelWithFormatter
	writer *syslog.Writer
}

// NewSysLogHandler instance
func NewSysLogHandler(priority syslog.Priority, tag string) (*SysLogHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSysLog handler instance with all custom options.
func NewSysLog(opt *SysLogOpt) (*SysLogHandler, error) { _ = "STUB: not implemented"; return nil, nil }

// init default log level

// Handle a log record
func (h *SysLogHandler) Handle(record *slog.Record) error { _ = "STUB: not implemented"; return nil }

// write log by level

// as info level

// Close handler
func (h *SysLogHandler) Close() error { _ = "STUB: not implemented"; return nil }

// Flush handler
func (h *SysLogHandler) Flush() error { _ = "STUB: not implemented"; return nil }
