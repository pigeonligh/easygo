package elog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
)

type Level int

const (
	InfoLevel Level = iota
	_
	_
	DebugLevel
)

type Logger struct {
	level  Level
	depth  int
	fields map[string]string
	outer  io.Writer
}

func (logger *Logger) V(level Level) *Logger {
	if logger == nil || level > logger.level {
		return nil
	}
	return logger
}

func (logger *Logger) With(fields map[string]string) *Logger {
	if logger == nil {
		return nil
	}
	newFields := make(map[string]string)
	for k, v := range logger.fields {
		nv, ok := fields[k]
		if ok && nv == "-" {
			continue
		}
		newFields[k] = v
	}
	for k, v := range fields {
		if v == "-" {
			continue
		}
		newFields[k] = v
	}
	return &Logger{
		level:  logger.level,
		depth:  logger.depth,
		fields: newFields,
		outer:  logger.outer,
	}
}

func (logger *Logger) output(msg string) {
	if logger == nil {
		return
	}
	_, file, line, ok := runtime.Caller(logger.depth + 1)
	if !ok {
		panic("source unknown")
	}

	fields := make(map[string]string)
	fields["source"] = fmt.Sprintf("%s:%d", file, line)
	for k, v := range logger.fields {
		fields[k] = v
	}
	fields["message"] = msg
	b, err := json.Marshal(fields)
	if err != nil {
		panic(err)
	}
	b = append(b, byte(','), byte('\n'))
	_, err = logger.outer.Write(b)
	if err != nil {
		panic(err)
	}
}

func (logger *Logger) print(v ...interface{}) {
	s := fmt.Sprint(v...)
	logger.depth++
	logger.output(s)
	logger.depth--
}

func (logger *Logger) printf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	logger.depth++
	logger.output(s)
	logger.depth--
}

// Info logs important message
func (logger *Logger) Info(v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Info",
	}).print(v...)
	logger.depth--
}

// Infof logs important message
func (logger *Logger) Infof(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Info",
	}).printf(format, v...)
	logger.depth--
}

// Warning logs warning message
func (logger *Logger) Warning(v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Warning",
	}).print(v...)
	logger.depth--
}

// Warningf logs warning message
func (logger *Logger) Warningf(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Warning",
	}).printf(format, v...)
	logger.depth--
}

// Error logs error message
func (logger *Logger) Error(v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Error",
	}).print(v...)
	logger.depth--
}

// Errorf logs error message
func (logger *Logger) Errorf(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Error",
	}).printf(format, v...)
	logger.depth--
}

// Fatal logs fatal message
func (logger *Logger) Fatal(v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Fatal",
	}).print(v...)
	logger.depth--
	os.Exit(255)
}

// Fatalf logs fatal message
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	if logger == nil {
		return
	}
	logger.depth++
	logger.With(map[string]string{
		"action": "Fatal",
	}).printf(format, v...)
	logger.depth--
	os.Exit(255)
}
