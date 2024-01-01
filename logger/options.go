package logger

import (
	"go.uber.org/zap"
)

type Options struct {
	appName        string
	encoding       Encoding
	samplingConfig *zap.SamplingConfig
	level          Level
}

type Option func(o *Options)

func WithEncoding(encoding Encoding) Option {
	return func(o *Options) {
		o.encoding = encoding
	}
}

func WithLevel(level Level) Option {
	return func(o *Options) {
		o.level = level
	}
}

func WithAppName(name string) Option {
	return func(o *Options) {
		o.appName = name
	}
}
