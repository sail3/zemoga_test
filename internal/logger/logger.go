package logger

import (
	"context"

	"go.uber.org/zap"
)

// CorrelationIDKey is the key to access the correlation id in the context.
const CorrelationIDKey string = "correlation_id"

const (
	errorKey   string = "error"
	serviceKey string = "service"
)

// Logger to print information to standard output.
type Logger interface {
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	With(key string, value interface{}) Logger
	WithError(err error) Logger
	WithCorrelation(ctx context.Context) Logger
}

// logger is an adapted zap logger.
type logger struct {
	*zap.SugaredLogger
	config *zap.Config
}

// With adds fields to the log.
func (l *logger) With(key string, value interface{}) Logger {
	return &logger{
		config:        l.config,
		SugaredLogger: l.SugaredLogger.With(key, value),
	}
}

// WithError adds an error to the log context.
func (l *logger) WithError(err error) Logger {
	return l.With(errorKey, err)
}

// WithCorrelation adds the correlation id to the log if it is found in the context.
func (l *logger) WithCorrelation(ctx context.Context) Logger {
	// add our correlation id if present.
	cid := ctx.Value(CorrelationIDKey)
	if cid == nil {
		return l
	}

	return l.With(CorrelationIDKey, cid)
}

// New creates a new logger.
func New(serviceName string, debug bool) Logger {
	if !debug {
		config := zap.NewProductionConfig()
		l, _ := config.Build()
		return &logger{l.Sugar().With(zap.String(serviceKey, serviceName)), &config}
	}

	config := zap.NewDevelopmentConfig()
	l, _ := config.Build()
	return &logger{l.Sugar().With(zap.String(serviceKey, serviceName)), &config}
}
