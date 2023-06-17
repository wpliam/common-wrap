package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/wpliam/common-wrap/client"
)

func init() {
	client.Register("redis", &Client{})
}

type Client struct {
}

func (c Client) Invoke(ctx context.Context, opts ...client.Option) error {
	o := &client.Options{}
	for _, opt := range opts {
		opt(o)
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     o.Target,
		Username: o.Username,
		Password: o.Password,
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	connPool[o.Name] = cli
	return nil
}
