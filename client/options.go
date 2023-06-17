package client

type Options struct {
	Name     string
	Target   string
	Username string
	Password string
	Protocol string
}

type Option func(opt *Options)

func WithName(name string) Option {
	return func(opt *Options) {
		opt.Name = name
	}
}

func WithTarget(target string) Option {
	return func(opt *Options) {
		opt.Target = target
	}
}

func WithUsername(username string) Option {
	return func(opt *Options) {
		opt.Username = username
	}
}

func WithPassword(password string) Option {
	return func(opt *Options) {
		opt.Password = password
	}
}

func WithProtocol(protocol string) Option {
	return func(opt *Options) {
		opt.Protocol = protocol
	}
}
