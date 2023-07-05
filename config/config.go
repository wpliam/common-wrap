package config

import (
	"github.com/wpliam/common-wrap/client"
	"os"

	"github.com/wpliam/common-wrap/plugin"
	"gopkg.in/yaml.v3"
)

var DefaultConfigPath = "./application.yaml"

type Config struct {
	Client  ClientConfig `yaml:"client"`
	Server  ServerConfig `yaml:"server"`
	Plugins plugin.Config
}

type ClientConfig struct {
	Service []client.ServiceConfig `yaml:"service"`
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
	content, err := os.ReadFile(DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err = yaml.Unmarshal(content, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
