package config

import (
	"os"
	"sync/atomic"

	"gopkg.in/yaml.v3"
)

const defaultConfigName = "application.yaml"

var cfg atomic.Value

func init() {
	cfg.Store(&AppConf{})
}

// AppConf 配置信息
type AppConf struct {
	Client *ClientConfig `yaml:"client"`
	Plugin *PluginConfig `yaml:"plugin"`
}

// LoadConfig 加载配置
func LoadConfig() error {
	bytes, err := os.ReadFile(defaultConfigName)
	if err != nil {
		return err
	}
	appConf := &AppConf{}
	if err = yaml.Unmarshal(bytes, appConf); err != nil {
		return err
	}
	cfg.Store(appConf)
	return nil
}

// GetConf 获取配置
func GetConf() *AppConf {
	return cfg.Load().(*AppConf)
}
