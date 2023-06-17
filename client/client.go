package client

import (
	"context"
	"sync"
)

var (
	clients = make(map[string]Client)
	lock    sync.RWMutex
)

type Client interface {
	Invoke(ctx context.Context, opts ...Option) error
}

func Register(name string, client Client) {
	lock.Lock()
	defer lock.Unlock()
	clients[name] = client
}

func Get(protocol string) Client {
	lock.RLock()
	c := clients[protocol]
	lock.RUnlock()
	return c
}
