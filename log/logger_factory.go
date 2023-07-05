package log

import (
	"github.com/wpliam/common-wrap/plugin"
	"sync"
)

const (
	defaultLogName = "default"
)

var (
	loggers = make(map[string]Logger)
	rw      sync.RWMutex
)

func init() {
	plugin.Register("log", &Factory{})
}

// Register 注册一个log
func Register(name string, logger Logger) {
	rw.Lock()
	defer rw.Unlock()
	loggers[name] = logger
}

// GetDefaultLogger 获取默认的log
func GetDefaultLogger() Logger {
	return Get(defaultLogName)
}

// Get 通过名称获取log
func Get(name string) Logger {
	rw.RLock()
	l := loggers[name]
	rw.RUnlock()
	return l
}

type Factory struct {
}

func (f *Factory) Setup(name string, decode plugin.Decoder) error {
	var cfg *LoggerConfig
	if err := decode.Decode(&cfg); err != nil {
		return err
	}
	Register(name, NewZapLog(cfg))
	return nil
}
