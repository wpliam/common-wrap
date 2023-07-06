package client

import (
	"github.com/wpliam/common-wrap/registry"
	"github.com/wpliam/common-wrap/selector"
)

type Client interface {
	Invoke(name string, opts ...Option) error
	Get() registry.Proxy
}

var New = func() Client {
	return &client{}
}

type client struct {
	proxy registry.Proxy
}

func (c *client) Invoke(name string, opts ...Option) error {
	o := getOptions(name)
	if o == nil {
		o = &Options{}
	}
	for _, opt := range opts {
		opt(o)
	}
	return c.selector(o)
}

func (c *client) selector(o *Options) error {
	opts := make([]selector.Option, 0)
	opts = append(opts,
		selector.WithTarget(o.Target),
		selector.WithUsername(o.Username),
		selector.WithPassword(o.Password),
	)
	proxy, err := o.selector.Select(opts...)
	if err != nil {
		return err
	}
	c.proxy = proxy
	return nil
}

func (c *client) Get() registry.Proxy {
	return c.proxy
}
