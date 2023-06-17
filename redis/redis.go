package redis

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	connPool = make(map[string]*redis.Client)
	lock     sync.RWMutex
)

var GetClientProxy = func(name string) *redis.Client {
	lock.RLock()
	defer lock.RUnlock()
	return connPool[name]
}
