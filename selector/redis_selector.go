package selector

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/wpliam/common-wrap/registry"
)

func init() {
	Register("redis", &redisSelector{})
}

type redisSelector struct {
}

func (r *redisSelector) Select(opt ...Option) (registry.Proxy, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     opts.Target,
		Username: opts.Username,
		Password: opts.Password,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return registry.NewRedisProxy(client), nil
}
