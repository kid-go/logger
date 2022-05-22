package formatter

import "time"

type Option func(*Options)

type Options struct {
	timeFormat string
	location   *time.Location
}

func WithTimeFormat(format string) Option {
	return func(o *Options) {
		o.timeFormat = format
	}
}

func WithLocation(location *time.Location) Option {
	return func(o *Options) {
		o.location = location
	}
}
