package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapSinLogger *zap.SugaredLogger
var once sync.Once

var zapLogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := zapLogger{cfg: cfg}
	logger.Init()
	return &logger
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	if level, ok := zapLogLevelMap[l.cfg.Logger.Level]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func (l *zapLogger) Init() {
	once.Do(
		func() {
			fileName := fmt.Sprintf("%s%s.log", l.cfg.Logger.FilePath, "logs")
			w := zapcore.AddSync(&lumberjack.Logger{
				Filename:   fileName,
				MaxSize:    100,
				MaxAge:     3,
				MaxBackups: 10,
				Compress:   true,
			})
			config := zap.NewProductionEncoderConfig()
			config.EncodeTime = zapcore.ISO8601TimeEncoder

			core := zapcore.NewCore(zapcore.NewJSONEncoder(config), w, l.getLogLevel())
			logger := zap.New(core, zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
			zapSinLogger = logger.With("AppName", "MyApp", "LoggerName", "ZapLogger")

		})
	l.logger = zapSinLogger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, err error, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Errorw(err.Error(), params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, err error, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Fatalw(err.Error(), params...)
	os.Exit(1)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKey(cat, sub, extra)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func prepareLogKey(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}
