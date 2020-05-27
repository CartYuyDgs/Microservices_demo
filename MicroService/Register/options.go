package Register

import "time"

type Options struct {
	Addrs   []string
	TimeOut time.Duration
}

type Option func(opts *Options)

func WithTimeOut(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.TimeOut = timeout
	}
}

func WithAddres(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}
