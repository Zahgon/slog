package handler

import (
	"github.com/gookit/slog"
)

// EmailOption struct
type EmailOption struct {
	SMTPHost string `json:"smtp_host"` // eg "smtp.gmail.com"
	SMTPPort int    `json:"smtp_port"` // eg 587
	FromAddr string `json:"from_addr"` // eg "yourEmail@gmail.com"
	Password string `json:"password"`
}

// EmailHandler struct
type EmailHandler struct {
	NopFlushClose
	slog.LevelWithFormatter
	// From the sender email information
	From EmailOption
	// ToAddresses email list
	ToAddresses []string
}

// NewEmailHandler instance
func NewEmailHandler(from EmailOption, toAddresses []string) *EmailHandler {
	_ = "STUB: not implemented"
	return nil
}

// to receivers

// init default log level

// Handle a log record
func (h *EmailHandler) Handle(r *slog.Record) error { _ = "STUB: not implemented"; return nil }
