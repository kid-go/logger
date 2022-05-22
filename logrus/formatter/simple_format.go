package formatter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type SimpleFormatter struct {
	options Options
}

const timeFormat = "2006-01-02 15:04:05"

func (s *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		t      = time.Now()
		format = timeFormat
	)
	if s.options.location != nil {
		t = time.Now().In(s.options.location)
	}
	if len(s.options.timeFormat) != 0 {
		format = s.options.timeFormat
	}
	msg := fmt.Sprintf("[%s] [%s] %s\n",
		t.Format(format),
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
	return []byte(msg), nil
}

func NewSimpleFormatter(options ...Option) *SimpleFormatter {
	formatter := &SimpleFormatter{options: Options{}}
	for _, o := range options {
		o(&formatter.options)
	}
	return formatter
}
