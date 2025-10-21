package logger

import (
	"fmt"
	"time"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

var minLevel string

func SetMinLevel(level string) {
	switch level {
	case DEBUG, INFO, WARNING, ERROR, FATAL:
		minLevel = level
	default:

	}
}

func Log(level string, message string, args ...interface{}) {
	if getLevelValue(level) < getLevelValue(minLevel) {
		return
	}
	timestamp := time.Now().UTC().Format(time.RFC3339)
	formattedMsg := fmt.Sprintf(message, args...)
	fmt.Printf("%s [%s] %s\n", timestamp, level, formattedMsg)
}

func getLevelValue(level string) int32 {
	switch level {
	case DEBUG:
		return 1
	case INFO:
		return 2
	case WARNING:
		return 3
	case ERROR:
		return 4
	case FATAL:
		return 5
	default:
		return 0
	}
}

func Debug(msg string, args ...interface{}) {
	Log(DEBUG, msg, args...)
}

func Info(msg string, args ...interface{}) {
	Log(INFO, msg, args...)
}

func Warning(msg string, args ...interface{}) {
	Log(WARNING, msg, args...)
}

func Error(msg string, args ...interface{}) {
	Log(ERROR, msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	Log(FATAL, msg, args...)
	panic(msg)
}
