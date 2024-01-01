package logger

import "log/slog"

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

// toSlog converts the type Level to slog.Level.
func (l *Level) toSlog() slog.Level {
	switch *l {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
