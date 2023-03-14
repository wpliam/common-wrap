package aredis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wpliap/common-wrap/config"
)

// NewRedisProxy 创建redis代理
func NewRedisProxy(name string) redis.UniversalClient {
	cfg := config.GetConnConf(name)
	if cfg == nil {
		panic(name + "redis conf not exist")
	}
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
	return client
}
