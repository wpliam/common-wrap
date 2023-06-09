package client

import "github.com/wpliam/common-wrap/selector"

type Options struct {
	Name     string
	Target   string
	Username string
	Password string
	Protocol string
	selector selector.Selector
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
		opt.selector = selector.Get(protocol)
	}
}
