package server

import (
	"fmt"
	stdhttp "github.com/wpliam/common-wrap/http"
	"net/http"
)

type Service interface {
	Listen() error
}

var NewService = func(name string, opts ...Option) Service {
	s := &service{
		name: name,
		opts: &Options{},
	}
	for _, opt := range opts {
		opt(s.opts)
	}
	return s
}

type service struct {
	name string
	opts *Options
}

func (s *service) Listen() error {
	switch s.opts.Protocol {
	case "http":
		svr := &http.Server{
			Addr:    fmt.Sprintf(":%d", s.opts.Port),
			Handler: stdhttp.GetHandler(s.name).Server(),
		}
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}
