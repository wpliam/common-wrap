package config

import (
	"github.com/wpliam/common-wrap/plugin"
	"gopkg.in/yaml.v3"
	"os"
)

const defaultConfigPath = "./application.yaml"

type Config struct {
	Client  ClientConfig `yaml:"client"`
	Server  ServerConfig `yaml:"server"`
	Plugins plugin.Config
}

type ClientConfig struct {
	Service []struct {
		Name     string `yaml:"name"`
		Target   string `yaml:"target"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Protocol string `yaml:"protocol"`
	} `yaml:"service"`
}

type ServerConfig struct {
	Service []struct {
		Name     string `yaml:"name"`
		Port     uint16 `yaml:"port"`
		Protocol string `yaml:"protocol"`
	} `yaml:"service"`
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	content, err := os.ReadFile(defaultConfigPath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err = yaml.Unmarshal(content, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
