package logger

import "github.com/Arshia-Izadyar/Jabama-clone/src/config"

type Logger interface {
	Init()
	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub SubCategory, err error, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub SubCategory, err error, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	logger := NewZapLogger(cfg)
	return logger
}
