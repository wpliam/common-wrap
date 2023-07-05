package selector

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func init() {
	Register("redis", &redisSelector{})
}

type redisSelector struct {
}

func (r *redisSelector) Select(opt ...Option) (interface{}, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     opts.Target,
		Username: opts.Username,
		Password: opts.Password,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return cli, nil
}
