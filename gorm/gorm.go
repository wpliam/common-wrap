package gorm

import (
	"gorm.io/gorm"
	"sync"
)

var (
	connPool = make(map[string]*gorm.DB)
	lock     sync.RWMutex
)

var GetClientProxy = func(name string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return connPool[name]
}
