package server

import (
	"fmt"
	"github.com/wpliam/common-wrap/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	services     map[string]Service
	shutdownHook []func()
	rw           sync.RWMutex
}

func (svr *Server) AddService(name string, service Service) {
	svr.rw.Lock()
	defer svr.rw.Unlock()
	if svr.services == nil {
		svr.services = make(map[string]Service)
	}
	svr.services[name] = service
}

func (svr *Server) AddShutdownHook(f func()) {
	svr.shutdownHook = append(svr.shutdownHook, f)
}

func (svr *Server) Run() {
	if len(svr.services) == 0 {
		panic("service is empty")
	}
	for name, service := range svr.services {
		go func(name string, service Service) {
			if err := service.Listen(); err != nil {
				panic(fmt.Sprintf("service [%s] listen err:%v", name, err))
			}
		}(name, service)
		log.Infof("service [%s] listen success process:%d", name, os.Getpid())
	}
	c := make(chan os.Signal, 0)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-c:
	}
	for _, hook := range svr.shutdownHook {
		hook()
	}
}
