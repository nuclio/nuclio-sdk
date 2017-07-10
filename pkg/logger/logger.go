package logger

// Logger is logging interface
type Logger interface {

	// emit a log entry of a given verbosity. the first argument may be an
	// object, a string or a format string. in case of the latter, the
	// following varargs are passed to a formatter (e.g. fmt.Sprintf)

	// Error logs an error message
	Error(format interface{}, vars ...interface{})
	// Warn logs a warning message
	Warn(format interface{}, vars ...interface{})
	// Info logs an info message
	Info(format interface{}, vars ...interface{})
	// Debug logs a debug message
	Debug(format interface{}, vars ...interface{})

	// emit a structured log entry. example:
	//
	// l.InfoWith("The message",
	// 	"first-key", "first-value",
	// 	"second-key", 2)
	//

	// ErrorWith emits an error message
	ErrorWith(format interface{}, vars ...interface{})
	// WarnWith emits a warning message
	WarnWith(format interface{}, vars ...interface{})
	// InfoWith emits an info message
	InfoWith(format interface{}, vars ...interface{})
	// DebugWith emits a debug message
	DebugWith(format interface{}, vars ...interface{})

	// Flush flushes buffered logs, if applicable
	Flush()

	// GetChild returns a child logger
	// If underlying logger supports hierarchal logging
	GetChild(name string) interface{}
}
