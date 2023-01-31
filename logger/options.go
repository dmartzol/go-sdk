package logger

import (
	"go.uber.org/zap/zapcore"
)

// Options represents logger options.
type Options struct {
	encoding     string
	levelEncoder zapcore.LevelEncoder
}

// Option represents a function that sets the options based on implementation.
type Option func(o *Options)

// WithFormat sets the specified format on the logger. Accepted values are "json" and "console".
func WithEncoding(encoding string) Option {
	if encoding != "console" && encoding != "json" {
		global.Fatalf("invalid log format '%s': must be text or json", encoding)
	}

	return func(o *Options) {
		o.encoding = encoding
	}
}

func WithColor() Option {
	return func(o *Options) {
		o.levelEncoder = zapcore.LowercaseColorLevelEncoder
	}
}
