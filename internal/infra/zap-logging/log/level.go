package log

import "strings"

// Level is a logger level.
type Level int8

// LevelKey is a logger level key.
const LevelKey = "level"

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = iota
	// LevelInfo is a logger info level.
	LevelInfo
	// LevelWarn is a logger warn level.
	LevelWarn
	// LevelError is a logger error level.
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return ""
	}
}

func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	}
	return LevelInfo
}
