package Register

import "time"

type Options struct {
	Addrs        []string
	TimeOut      time.Duration
	RegistryPath []string
	HeartBeat    int64
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

func WithRegistryPath(path []string) Option {
	return func(opts *Options) {
		opts.RegistryPath = path
	}
}

func WithRegistryHeartBeat(beat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = beat
	}
}
