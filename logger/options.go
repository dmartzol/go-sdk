package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options represents logger options.
type Options struct {
	encoding       string
	levelEncoder   zapcore.LevelEncoder
	samplingConfig *zap.SamplingConfig
}

type Option func(o *Options)

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

func WithSampling(initial, thereafter int) Option {
	return func(o *Options) {
		o.samplingConfig = &zap.SamplingConfig{
			Initial:    initial,
			Thereafter: thereafter,
		}
	}
}
