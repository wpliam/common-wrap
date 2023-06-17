package log

import (
	"github.com/wpliap/common-wrap/plugin"
	"sync"
)

const (
	defaultLogName = "default"
)

var (
	DefaultLogger Logger
	loggers       = make(map[string]Logger)
	rw            sync.RWMutex
)

func init() {
	plugin.Register("log", &Factory{})
}

// Register 注册一个log
func Register(name string, logger Logger) {
	rw.Lock()
	defer rw.Unlock()
	if _, ok := loggers[name]; ok && name != defaultLogName {
		panic("register name exist " + name)
	}
	if name == defaultLogName {
		DefaultLogger = logger
	}
	loggers[name] = logger
}

// GetDefaultLogger 获取默认的log
func GetDefaultLogger() Logger {
	rw.RLock()
	l := DefaultLogger
	rw.RLocker()
	return l
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
