package rotate

import (
	"github.com/leor-w/kid/plugin/logger/logrus/formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

func NewSimpleRotate(opts ...Option) (logrus.Hook, error) {
	options := Options{
		path:          "./logs",
		logPrefix:     "kid_",
		logName:       "%Y%m%d",
		logSuffix:     ".log",
		link:          "",
		rotate:        time.Hour * 24,
		maxAge:        time.Hour * 24 * 15,
		location:      "Asia/ShangHai",
		logTimeFormat: "2006-01-02 15:04:05",
	}
	for _, o := range opts {
		o(&options)
	}
	writer, err := rotatelogs.New(
		path.Join(options.path, options.logPrefix+options.logName+options.logSuffix),
		rotatelogs.WithLinkName(options.link),
		rotatelogs.WithRotationTime(options.rotate),
		rotatelogs.WithMaxAge(options.maxAge),
	)
	if err != nil {
		return nil, err
	}
	location, _ := time.LoadLocation(options.location)
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, formatter.NewSimpleFormatter(
		formatter.WithTimeFormat(options.logTimeFormat),
		formatter.WithLocation(location),
	)), nil
}
