package log

import "sync"

const (
	defaultLogName  = "default"
	defaultFileName = "common.log"
)

var (
	DefaultLogger Logger
	loggers       = make(map[string]Logger)
	rw            sync.RWMutex
)

func init() {
	Register(defaultLogName, NewZapLog(defaultFileName))
}

func Register(name string, logger Logger) {
	rw.Lock()
	defer rw.Unlock()
	if _, ok := loggers[name]; ok {
		return
	}
	if name == defaultLogName {
		DefaultLogger = logger
	}
	loggers[name] = logger
}

func GetDefaultLogger() Logger {
	rw.RLock()
	l := DefaultLogger
	rw.RLocker()
	return l
}

func Get(name string) Logger {
	rw.RLock()
	l := loggers[name]
	rw.RUnlock()
	return l
}
