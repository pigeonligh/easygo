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
		depth:  1,
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
	defaultLogger.Info(v...)
}

// Infof logs important message
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Warning logs warning message
func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

// Warningf logs warning message
func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

// Error logs error message
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Errorf logs error message
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Fatal logs fatal message
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

// Fatalf logs fatal message
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}
