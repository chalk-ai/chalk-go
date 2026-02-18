package chalk

import (
	"fmt"
	"io"
	"os"
)

// StdOutLeveledLogger is a leveled logger implementation.
//
// It prints warnings and errors to `os.Stderr` and other messages to
// `os.Stdout`.
type StdOutLeveledLogger struct {
	// Level is the minimum logging level that will be emitted by this logger.
	//
	// For example, a Level set to LevelWarn will emit warnings and errors, but
	// not informational or debug messages.
	//
	// Always set this with a constant like LevelWarn because the individual
	// values are not guaranteed to be stable.
	Level Level

	// Internal testing use only.
	stderrOverride io.Writer
	stdoutOverride io.Writer
}

// Debugf logs a debug message using Printf conventions.
func (l *StdOutLeveledLogger) Debugf(format string, v ...any) {
	if l.Level >= LevelDebug {
		fmt.Fprintf(l.stdout(), "[DEBUG] "+format+"\n", v...)
	}
}

// Errorf logs a warning message using Printf conventions.
func (l *StdOutLeveledLogger) Errorf(format string, v ...any) {
	// Infof logs a debug message using Printf conventions.
	if l.Level >= LevelError {
		fmt.Fprintf(l.stderr(), "[ERROR] "+format+"\n", v...)
	}
}

// Infof logs an informational message using Printf conventions.
func (l *StdOutLeveledLogger) Infof(format string, v ...any) {
	if l.Level >= LevelInfo {
		fmt.Fprintf(l.stdout(), "[INFO] "+format+"\n", v...)
	}
}

// Warnf logs a warning message using Printf conventions.
func (l *StdOutLeveledLogger) Warnf(format string, v ...any) {
	if l.Level >= LevelWarn {
		fmt.Fprintf(l.stderr(), "[WARN] "+format+"\n", v...)
	}
}

func (l *StdOutLeveledLogger) stderr() io.Writer {
	if l.stderrOverride != nil {
		return l.stderrOverride
	}

	return os.Stderr
}

func (l *StdOutLeveledLogger) stdout() io.Writer {
	if l.stdoutOverride != nil {
		return l.stdoutOverride
	}

	return os.Stdout
}
