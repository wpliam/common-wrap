package selector

type Options struct {
	Target   string
	Username string
	Password string
}

type Option func(*Options)

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
