package server

type Options struct {
	Port     uint16
	Protocol string
}

type Option func(opt *Options)

func WithPort(port uint16) Option {
	return func(opt *Options) {
		opt.Port = port
	}
}

func WithProtocol(protocol string) Option {
	return func(opt *Options) {
		opt.Protocol = protocol
	}
}
