package logrus

import (
	"github.com/leor-w/kid/logger"
	"github.com/sirupsen/logrus"
)

type Options struct {
	logger.Options
	formatter  logrus.Formatter
	reportCall bool
	exitFunc   func(int)
}

type formatterKey struct{}

func WithFormatter(formatter logrus.Formatter) logger.Option {
	return logger.SetOption(formatterKey{}, formatter)
}

type reportCallKey struct{}

func WithReportCall(reportCall bool) logger.Option {
	return logger.SetOption(reportCallKey{}, reportCall)
}

type exitFuncKey struct{}

func WithExitFunc(exec func(int)) logger.Option {
	return logger.SetOption(exitFuncKey{}, exec)
}
