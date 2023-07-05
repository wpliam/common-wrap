package server

import (
	"fmt"
	"net/http"

	jhttp "github.com/wpliam/common-wrap/http"
)

type Service interface {
	Listen() error
}

var New = func(name string, opts ...Option) Service {
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
		handler := jhttp.GetHandler(s.name)
		if handler == nil {
			return nil
		}
		svr := &http.Server{
			Addr:    fmt.Sprintf(":%d", s.opts.Port),
			Handler: handler.Server(),
		}
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}
