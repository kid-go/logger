package rotate

import "time"

type Option func(*Options)

type Options struct {
	// path 日志文件存储路径
	path string
	// logPrefix 日志文件前缀
	logPrefix string
	// logName 日志文件名称
	logName string
	// logSuffix 日志文件后缀
	logSuffix string
	// link 日志文件软连接到的位置
	link string
	// rotate 日志分割时间
	rotate time.Duration
	// maxAge 分割日志保存最长时间
	maxAge time.Duration
	// location 日志文件时间使用时区
	location string
	// logTimeFormat 日志时间格式化模板
	logTimeFormat string
}

func WithPath(path string) Option {
	return func(o *Options) {
		o.path = path
	}
}

func WithLogPrefix(prefix string) Option {
	return func(o *Options) {
		o.logPrefix = prefix
	}
}

func WithLogName(logName string) Option {
	return func(o *Options) {
		o.logName = logName
	}
}

func WithLogSuffix(suffix string) Option {
	return func(o *Options) {
		o.logSuffix = suffix
	}
}

func WithLink(link string) Option {
	return func(o *Options) {
		o.link = link
	}
}

func WithRotate(rotate time.Duration) Option {
	return func(o *Options) {
		o.rotate = rotate
	}
}

func WithMaxAge(maxAge time.Duration) Option {
	return func(o *Options) {
		o.maxAge = maxAge
	}
}

func WithLocation(location string) Option {
	return func(o *Options) {
		o.location = location
	}
}

func WithLogTimeFormat(timeFormat string) Option {
	return func(o *Options) {
		o.logTimeFormat = timeFormat
	}
}
