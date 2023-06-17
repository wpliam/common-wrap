package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/wpliap/common-wrap/client"
	"github.com/wpliap/common-wrap/config"
)

type Server struct {
	services     map[string]Service
	shutdownHook []func()
	rw           sync.RWMutex
}

func New() *Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("load config err " + err.Error())
	}
	if cfg.Plugins != nil {
		if err = cfg.Plugins.Setup(); err != nil {
			panic("plugin setup err" + err.Error())
		}
	}

	for _, s := range cfg.Client.Service {
		cli := client.Get(s.Protocol)
		opts := []client.Option{
			client.WithName(s.Name),
			client.WithTarget(s.Target),
			client.WithUsername(s.Username),
			client.WithPassword(s.Password),
			client.WithProtocol(s.Protocol),
		}
		if err = cli.Invoke(context.Background(), opts...); err != nil {
			panic(fmt.Sprintf("client [%s] invoke err:%v", s.Name, err))
		}
	}

	svr := &Server{}
	for _, s := range cfg.Server.Service {
		svr.AddService(s.Name, NewService(s.Name, WithPort(s.Port), WithProtocol(s.Protocol)))
	}
	return svr
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
