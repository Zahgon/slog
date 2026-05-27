package slog

import (
	"runtime"
)

//
// Processor interface
//

// Processor interface definition
type Processor interface {
	// Process record
	Process(record *Record)
}

// ProcessorFunc wrapper definition
type ProcessorFunc func(record *Record)

// Process record
func (fn ProcessorFunc) Process(record *Record) {
	_ = "STUB: not implemented"

	// ProcessableHandler interface
	return
}

type ProcessableHandler interface {
	// AddProcessor add a processor
	AddProcessor(Processor)
	// ProcessRecord handle a record
	ProcessRecord(record *Record)
}

// Processable definition
type Processable struct {
	processors []Processor
}

// AddProcessor to the handler
func (p *Processable) AddProcessor(processor Processor) { _ = "STUB: not implemented"; return }

// ProcessRecord process record
func (p *Processable) ProcessRecord(r *Record) {
	_ = "STUB: not implemented"
	// processing log record
	return
}

//
// there are some built-in processors
//

// AddHostname to record
func AddHostname() Processor { _ = "STUB: not implemented"; return *new(Processor) }

// AddUniqueID to record
func AddUniqueID(fieldName string) Processor { _ = "STUB: not implemented"; return *new(Processor) }

// MemoryUsage get memory usage.
var MemoryUsage ProcessorFunc = func(record *Record) {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	record.SetExtraValue("memoryUsage", stat.Alloc)
}

// AppendCtxKeys append context keys to Record.Fields
func AppendCtxKeys(keys ...string) Processor { _ = "STUB: not implemented"; return *new(Processor) }

// CtxKeysProcessor append context keys to Record.Data, Record.Fields, Record.Extra
//   - dist: "data" | "fields" | "extra"
func CtxKeysProcessor(dist string, keys ...string) Processor {
	_ = "STUB: not implemented"
	return *new(Processor)
}
