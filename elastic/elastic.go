package elastic

import (
	"github.com/olivere/elastic/v7"
	"sync"
)

var (
	connPool = make(map[string]*elastic.Client)
	lock     sync.RWMutex
)

var GetClientProxy = func(name string) *elastic.Client {
	lock.RLock()
	defer lock.RUnlock()
	return connPool[name]
}
