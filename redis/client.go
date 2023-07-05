package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/wpliam/common-wrap/client"
)

var NewClientProxy = func(name string, opts ...client.Option) *redis.Client {
	options := make([]client.Option, 0, len(opts)+1)
	options = append(options, opts...)
	options = append(options,
		client.WithProtocol("redis"),
	)
	newClient := client.New()
	if err := newClient.Invoke(name, options...); err != nil {
		panic("redis client invoke err" + err.Error())
	}
	proxy := newClient.Get()
	if proxy == nil {
		panic("get redis client proxy not exist")
	}
	return proxy.(*redis.Client)
}
