package logger

import (
	"github.com/leor-w/kid/logger"
	"github.com/leor-w/kid/plugin/logger/logrus"
	"github.com/leor-w/kid/plugin/logger/logrus/formatter"
	"sync"
	"time"
)

var (
	_defaultLogger *Logger
	_defaultOnce   sync.Once
)

type Logger struct {
	logger.Logger
}

func (log *Logger) Init(map[string]interface{}) error {
	NewLogger(nil)
	return nil
}

func NewLogger(logger *Logger) *Logger {
	if logger != nil {
		_defaultLogger = logger
	}
	if _defaultLogger == nil {
		_defaultOnce.Do(func() {
			_defaultLogger = &Logger{Logger: newDefaultLogger()}
		})
	}
	return _defaultLogger
}

func newDefaultLogger() logger.Logger {
	location, _ := time.LoadLocation("Asia/ShangHai")
	return logrus.NewLogger(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFields(map[string]interface{}{
			"hello": "world",
		}),
		logrus.WithReportCall(true),
		logrus.WithFormatter(formatter.NewSimpleFormatter(
			formatter.WithLocation(location))),
	)
}

func AddHook(hook logger.Hook) {
	_defaultLogger.Hook(hook)
}

func AddWorker(worker logger.Worker) {
	AddHook(_defaultLogger.NewKidHook(worker))
}

func Trace(args ...interface{}) {
	_defaultLogger.Log(logger.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	_defaultLogger.Log(logger.DebugLevel, args...)
}

func Info(args ...interface{}) {
	_defaultLogger.Log(logger.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	_defaultLogger.Log(logger.WarnLevel, args...)
}

func Error(args ...interface{}) {
	_defaultLogger.Log(logger.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	_defaultLogger.Log(logger.FatalLevel, args...)
}

func Tracef(fmt string, args ...interface{}) {
	_defaultLogger.Logf(logger.TraceLevel, fmt, args...)
}

func Debugf(fmt string, args ...interface{}) {
	_defaultLogger.Logf(logger.DebugLevel, fmt, args...)
}

func Infof(fmt string, args ...interface{}) {
	_defaultLogger.Logf(logger.InfoLevel, fmt, args...)
}

func Warnf(fmt string, args ...interface{}) {
	_defaultLogger.Logf(logger.WarnLevel, fmt, args...)
}

func Errorf(fmt string, args ...interface{}) {
	_defaultLogger.Logf(logger.ErrorLevel, fmt, args...)
}

func Fatalf(fmt string, args ...interface{}) {
	_defaultLogger.Logf(logger.FatalLevel, fmt, args...)
}
