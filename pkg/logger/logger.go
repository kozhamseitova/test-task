package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// "go.uber.org/zap/zapcore"
)

type Logger interface {
	Errorf(c context.Context, format string, args ...interface{})
	Fatalf(c context.Context, format string, args ...interface{})
	Infof(c context.Context, format string, args ...interface{})
	Warnf(c context.Context, format string, args ...interface{})
	Debugf(c context.Context, format string, args ...interface{})
}

type logger struct {
	zapLogger *zap.SugaredLogger
}

func NewLogger(production bool) (Logger, error) {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	// config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return &logger{
		zapLogger: l.Sugar(),
	}, nil
}

func (l *logger) Errorf(c context.Context, format string, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Errorf(format, args)
}
func (l *logger) Fatalf(c context.Context, format string, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Fatalf(format, args)
}
func (l *logger) Fatal(c context.Context, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Fatal(args)
}
func (l *logger) Infof(c context.Context, format string, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Infof(format, args)
}
func (l *logger) Warnf(c context.Context, format string, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Warnf(format, args)
}
func (l *logger) Debugf(c context.Context, format string, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Debugf(format, args)
}
func (l *logger) Printf(c context.Context, format string, args ...interface{}) {
	l.zapLogger.With("traceID", c.Value("traceID")).Infof(format, args)
}
