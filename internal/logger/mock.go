package logger

import "context"

// mock to logger.
type mock struct{}

func (m mock) Error(_ ...interface{}) {
}

func (m mock) Errorf(_ string, _ ...interface{}) {
}

func (m mock) Debug(_ ...interface{}) {
}

func (m mock) Debugf(_ string, _ ...interface{}) {
}

func (m mock) Info(_ ...interface{}) {
}

func (m mock) Infof(_ string, _ ...interface{}) {
}

func (m mock) Warn(_ ...interface{}) {
}

func (m mock) Warnf(_ string, _ ...interface{}) {
}

func (m mock) With(_ string, _ interface{}) Logger {
	return m
}

func (m mock) WithError(_ error) Logger {
	return m
}

func (m mock) WithCorrelation(_ context.Context) Logger {
	return m
}

// Mock returns a new logger mock.
func Mock() Logger {
	return mock{}
}
