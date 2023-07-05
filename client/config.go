package client

import "sync"

var (
	options = make(map[string]*Options)
	lock    sync.RWMutex
)

func getOptions(name string) *Options {
	lock.RLock()
	opt := options[name]
	lock.RUnlock()
	return opt
}

type ServiceConfig struct {
	Name     string `yaml:"name"`
	Target   string `yaml:"target"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
}

func (c *ServiceConfig) genOptions() []Option {
	opts := make([]Option, 0)
	opts = append(opts,
		WithProtocol(c.Protocol),
		WithName(c.Name),
		WithTarget(c.Target),
		WithUsername(c.Username),
		WithPassword(c.Password),
	)
	return opts
}

func RegisterClientConfig(name string, cfg *ServiceConfig) {
	opts := cfg.genOptions()
	lock.Lock()
	defer lock.Unlock()
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	options[name] = opt
}
