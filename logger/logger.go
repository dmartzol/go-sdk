package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

var global *slog.Logger

func init() {
	global = New()
}

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func NewWithOptions(opts ...Option) *slog.Logger {
	// default options
	options := &Options{
		encoding:       EncodingJSON,
		samplingConfig: nil,
	}

	for _, o := range opts {
		o(options)
	}

	handlerOptions := &slog.HandlerOptions{
		Level: options.level.toSlog(),
	}

	var logger *slog.Logger
	switch options.encoding {
	case EncodingJSON:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, handlerOptions))
	case EncodingConsole:
		tintOptions := &tint.Options{
			Level: options.level.toSlog(),
		}
		logger = slog.New(tint.NewHandler(os.Stdout, tintOptions))
	default:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, handlerOptions))
	}

	return logger
}

func New() *slog.Logger {
	return NewWithOptions()
}
