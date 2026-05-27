package slog

//
// ---------------------------------------------------------------------------
// Do write log message
// ---------------------------------------------------------------------------
//

// func (r *Record) logWrite(level Level) {
// Will reduce memory allocation once
// r.Message = strutil.Byte2str(message)

// var buf *bytes.Buffer
// buf = bufferPool.Get().(*bytes.Buffer)
// defer bufferPool.Put(buf)
// r.Buffer = buf

// TODO release on here ??
// defer r.logger.releaseRecord(r)
// r.logger.writeRecord(level, r)
// r.Buffer = nil
// }

// Init something for record(eg: time, level name).
func (r *Record) Init(lowerLevelName bool) {
	_ = "STUB: not implemented"

	// use lower level name
	return
}

// init log time

// r.microSecond = r.Time.Nanosecond() / 1000

// Init something for record.
func (r *Record) beforeHandle(l *Logger) {
	_ = "STUB: not implemented"
	// log caller. will alloc 3 times
	return
}

// processing log record

// do write record to handlers, will add lock.
func (l *Logger) writeRecord(level Level, r *Record) { _ = "STUB: not implemented"; return }

// reset init flag, useful for repeat use Record

// init record, call processors

// do write a log message by handler

// ---- after write log ----

// flush logs on level <= error level.

// has been in lock
