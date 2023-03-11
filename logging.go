package chalk

// Level represents a logging level.
type Level uint32

const (
	// LevelNull sets a logger to show no messages at all.
	LevelNull Level = 0

	// LevelError sets a logger to show error messages only.
	LevelError Level = 1

	// LevelWarn sets a logger to show warning messages or anything more
	// severe.
	LevelWarn Level = 2

	// LevelInfo sets a logger to show informational messages or anything more
	// severe.
	LevelInfo Level = 3

	// LevelDebug sets a logger to show informational messages or anything more
	// severe.
	LevelDebug Level = 4
)

// DefaultLeveledLogger is the default logger that the library will use to log
// errors, warnings, and informational messages.
//
// LeveledLogger is implemented by StdOutLeveledLogger, and one can be
// initialized at the desired level of logging.  LeveledLogger also
// provides out-of-the-box compatibility with a Logrus Logger, but may require
// a thin shim for use with other logging libraries that use less standard
// conventions like Zap.
//
// This Logger will be inherited by any backends created by default, but will
// be overridden if a backend is created with GetBackendWithConfig with a
// custom StdOutLeveledLogger set.
var DefaultLeveledLogger LeveledLogger = &StdOutLeveledLogger{
	Level: LevelError,
}

// LeveledLogger provides a basic leveled logging interface for
// printing debug, informational, warning, and error messages.
//
// It's implemented by StdOutLeveledLogger and also provides out-of-the-box
// compatibility with a Logrus Logger, but may require a thin shim for use with
// other logging libraries that you use less standard conventions like Zap.
type LeveledLogger interface {
	// Debugf logs a debug message using Printf conventions.
	Debugf(format string, v ...interface{})

	// Errorf logs a warning message using Printf conventions.
	Errorf(format string, v ...interface{})

	// Infof logs an informational message using Printf conventions.
	Infof(format string, v ...interface{})

	// Warnf logs a warning message using Printf conventions.
	Warnf(format string, v ...interface{})
}
