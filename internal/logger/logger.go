package logger

import (
	"log"
	"os"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	level LogLevel
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

// Log is the exported, initialized logger instance
var Log *Logger

// init function initializes Log with the log level from LOG_LEVEL environment variable
func init() {
	level := parseLogLevelFromEnv()
	Log = NewLogger(level)
}

// parseLogLevelFromEnv reads the LOG_LEVEL environment variable and returns the corresponding LogLevel.
// Defaults to INFO if LOG_LEVEL is unset or invalid.
func parseLogLevelFromEnv() LogLevel {
	logLevelStr := os.Getenv("LOG_LEVEL")
	switch strings.ToUpper(logLevelStr) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return INFO // Default to INFO if LOG_LEVEL is not set or invalid
	}
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
		debug: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warn:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatal: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Debug logs debug messages if the level is set to DEBUG or lower
func (l *Logger) Debug(msg string, v ...interface{}) {
	if l.level <= DEBUG {
		l.debug.Printf(msg, v...)
	}
}

// Info logs informational messages if the level is set to INFO or lower
func (l *Logger) Info(msg string, v ...interface{}) {
	if l.level <= INFO {
		l.info.Printf(msg, v...)
	}
}

// Warn logs warning messages if the level is set to WARN or lower
func (l *Logger) Warn(msg string, v ...interface{}) {
	if l.level <= WARN {
		l.warn.Printf(msg, v...)
	}
}

// Error logs error messages if the level is set to ERROR or lower
func (l *Logger) Error(msg string, v ...interface{}) {
	if l.level <= ERROR {
		l.error.Printf(msg, v...)
	}
}

// Fatal logs fatal messages if the level is set to FATAL or lower
func (l *Logger) Fatal(msg string, v ...interface{}) {
	if l.level <= FATAL {
		l.fatal.Printf(msg, v...)
	}
}

// Wrapper functions to simplify logging

func Debug(msg string, v ...interface{}) {
	Log.Debug(msg, v...)
}

func Info(msg string, v ...interface{}) {
	Log.Info(msg, v...)
}

func Warn(msg string, v ...interface{}) {
	Log.Warn(msg, v...)
}

func Error(msg string, v ...interface{}) {
	Log.Error(msg, v...)
}

func Fatal(msg string, v ...interface{}) {
	Log.Fatal(msg, v...)
}
