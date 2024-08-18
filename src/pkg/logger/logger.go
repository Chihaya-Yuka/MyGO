package logger

import (
	"log"
	"os"
)

// Logger represents a logger.
type Logger struct {
	*log.Logger
}

// NewLogger returns a new logger.
func NewLogger(level string) (*Logger, error) {
	var logger *log.Logger
	switch level {
	case "debug":
		logger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "info":
		logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "warn":
		logger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	case "error":
		logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	default:
		return nil, fmt.Errorf("invalid log level: %s", level)
	}
	return &Logger{logger}, nil
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string) {
	l.Logger.Println(msg)
}

// Info logs an info message.
func (l *Logger) Info(msg string) {
	l.Logger.Println(msg)
}

// Warn logs a warn message.
func (l *Logger) Warn(msg string) {
	l.Logger.Println(msg)
}

// Error logs an error message.
func (l *Logger) Error(msg string) {
	l.Logger.Println(msg)
}
