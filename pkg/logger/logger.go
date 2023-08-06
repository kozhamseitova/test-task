package logger

import (
	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type logger struct {
	zapLogger *zap.SugaredLogger
}

func NewLogger(production bool) (Logger, error) {
	config := zap.NewProductionConfig()
	config.DisableStacktrace = true
	// config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return &logger{
		zapLogger: l.Sugar(),
	}, nil
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.zapLogger.Errorf(format, args)
}
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.zapLogger.Fatalf(format, args)
}
func (l *logger) Fatal(args ...interface{}) {
	l.zapLogger.Fatal(args)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.zapLogger.Infof(format, args)
}
func (l *logger) Warnf(format string, args ...interface{}) {
	l.zapLogger.Warnf(format, args)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.zapLogger.Debugf(format, args)
}
func (l *logger) Printf(format string, args ...interface{}) {
	l.zapLogger.Infof(format, args)
}
func (l *logger) Println(args ...interface{}) {
	l.zapLogger.Info(args, "\n")
}
