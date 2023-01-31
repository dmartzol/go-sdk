package logger

import "fmt"

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

func New(structuredLogging bool) (Logger, error) {
	if structuredLogging {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("unable to create production logger: %w", err)
		}
		return logger.Sugar(), nil
	} else {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err := config.Build()
		if err != nil {
			return nil, fmt.Errorf("unable to create development logger: %w", err)
		}
		return logger.Sugar(), nil
	}
}
