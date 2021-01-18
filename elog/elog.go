package elog

import (
	"io"
	"os"
)

var (
	defaultLogger *Logger = nil
)

func Default() {
	Setup(InfoLevel, os.Stdout)
}

func Debug() {
	Setup(DebugLevel, os.Stdout)
}

func Setup(level Level, outer io.Writer) {
	defaultLogger = &Logger{
		level:  level,
		depth:  0,
		fields: map[string]string{},
		outer:  outer,
	}
}

func V(level Level) *Logger {
	return defaultLogger.V(level)
}

func With(fields map[string]string) *Logger {
	return defaultLogger.With(fields)
}

// Info logs important message
func Info(v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Info(v...)
	defaultLogger.depth--
}

// Infof logs important message
func Infof(format string, v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Infof(format, v...)
	defaultLogger.depth--
}

// Warning logs warning message
func Warning(v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Warning(v...)
	defaultLogger.depth--
}

// Warningf logs warning message
func Warningf(format string, v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Warningf(format, v...)
	defaultLogger.depth--
}

// Error logs error message
func Error(v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Error(v...)
	defaultLogger.depth--
}

// Errorf logs error message
func Errorf(format string, v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Errorf(format, v...)
	defaultLogger.depth--
}

// Fatal logs fatal message
func Fatal(v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Fatal(v...)
	defaultLogger.depth--
}

// Fatalf logs fatal message
func Fatalf(format string, v ...interface{}) {
	defaultLogger.depth++
	defaultLogger.Fatalf(format, v...)
	defaultLogger.depth--
}
