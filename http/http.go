package http

import (
	"net/http"
	"sync"
)

var (
	handlers = make(map[string]Handler)
	lock     sync.RWMutex
)

type Handler interface {
	Server() http.Handler
}

func Register(name string, h Handler) {
	lock.Lock()
	defer lock.Unlock()
	handlers[name] = h
}

func GetHandler(name string) Handler {
	lock.RLock()
	h := handlers[name]
	lock.RUnlock()
	return h
}
