package jsonlogger

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

func New(output io.Writer, minLevel Level) *Logger {
	return &Logger{
		minLevel: minLevel,
		output:   output,
	}
}

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	output   io.Writer
	minLevel Level
	mu       sync.Mutex
}

func (l *Logger) LogInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

func (l *Logger) LogError(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

func (l *Logger) LogFatal(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	aux := struct {
		Level      string
		Time       string
		Message    string
		Properties map[string]string
		Trace      string
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	var line []byte
	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + " unable to marshal log message " + err.Error())
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	return l.output.Write(line)
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}
