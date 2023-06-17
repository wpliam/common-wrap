package plugin

import (
	"gopkg.in/yaml.v3"
	"sync"
)

var (
	plugins = make(map[string]Factory)
	rw      sync.RWMutex
)

func Register(name string, factory Factory) {
	rw.Lock()
	defer rw.Unlock()
	plugins[name] = factory
}

func Get(name string) Factory {
	rw.RLock()
	f := plugins[name]
	rw.RUnlock()
	return f
}

type Config map[string]map[string]yaml.Node

func (c Config) Setup() error {
	for name, factory := range c {
		f := Get(name)
		for key, node := range factory {
			decode := &YamlNodeDecode{Node: &node}
			if err := f.Setup(key, decode); err != nil {
				return err
			}
		}
	}
	return nil
}

type Decoder interface {
	Decode(cfg interface{}) error
}

type YamlNodeDecode struct {
	Node *yaml.Node
}

func (d *YamlNodeDecode) Decode(cfg interface{}) error {
	return d.Node.Decode(cfg)
}

type Factory interface {
	Setup(name string, decode Decoder) error
}
