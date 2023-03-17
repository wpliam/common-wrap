package redis

import (
	"context"
	"fmt"
	"github.com/wpliap/common-wrap/log"

	"github.com/go-redis/redis/v8"
	"github.com/wpliap/common-wrap/config"
)

// NewRedisProxy 创建redis代理
func NewRedisProxy(name string, opt ...config.Option) redis.UniversalClient {
	cfg := config.GetConnConf(name, opt...)
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{fmt.Sprintf("%s:%d", cfg.GetHost(), cfg.GetPort())},
		Username: cfg.GetUsername(),
		Password: cfg.GetPassword(),
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("redis client ping err:" + err.Error())
	}
	log.Infof("create redis proxy success name:%s host:%s port:%d", name, cfg.GetHost(), cfg.GetPort())
	return client
}
