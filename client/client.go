package client

import "github.com/wpliam/common-wrap/selector"

type Client interface {
	Invoke(name string, opts ...Option) error
	Get() interface{}
}

var New = func() Client {
	return &client{}
}

type client struct {
	proxy interface{}
}

func (c *client) Invoke(name string, opts ...Option) error {
	opt := getOptions(name)
	if opt == nil {
		opt = &Options{}
	}
	for _, o := range opts {
		o(opt)
	}
	return c.selector(opt)
}

func (c *client) selector(opt *Options) error {
	opts := make([]selector.Option, 0)
	opts = append(opts,
		selector.WithTarget(opt.Target),
		selector.WithUsername(opt.Username),
		selector.WithPassword(opt.Password),
	)
	proxy, err := opt.selector.Select(opts...)
	if err != nil {
		return err
	}
	c.proxy = proxy
	return nil
}

func (c *client) Get() interface{} {
	return c.proxy
}
