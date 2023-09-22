package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Infof(template string, args ...interface{})
	Info(args ...interface{})
	Errorf(template string, args ...interface{})
	Error(args ...interface{})
	Sync() error
}

func New() (Logger, error) {
	inner, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return inner.Sugar(), nil
}
