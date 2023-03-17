package config

// PluginConfig 插件配置
type PluginConfig struct {
	Log map[string]*LogConfig `yaml:"log"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogPath    string `yaml:"log_path"`
	Filename   string `yaml:"filename"`
	MaxAge     int    `yaml:"max_age"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}

// GetLogConf 获取配置信息
func GetLogConf(name string) *LogConfig {
	if name == "default" {
		return defaultLogConfig()
	}
	return GetConf().Plugin.Log[name]
}

func defaultLogConfig() *LogConfig {
	return &LogConfig{
		LogPath:    "/usr/local/service/log",
		Filename:   "flow.log",
		MaxAge:     7,
		MaxSize:    10,
		MaxBackups: 10,
		Compress:   false,
	}
}

// GetLogPath 获取日志目录
func (l *LogConfig) GetLogPath() string {
	if l == nil || l.LogPath == "" {
		return "/usr/local/service/log"
	}
	return l.LogPath
}

// GetFilename 获取文件名
func (l *LogConfig) GetFilename() string {
	if l == nil || l.Filename == "" {
		return "flow.log"
	}
	return l.Filename
}

// GetMaxAge 最大日志保留天数
func (l *LogConfig) GetMaxAge() int {
	if l == nil || l.MaxAge == 0 {
		return 7
	}
	return l.MaxAge
}

// GetMaxSize 本地文件滚动日志大小,单位MB
func (l *LogConfig) GetMaxSize() int {
	if l == nil || l.MaxSize == 0 {
		return 10
	}
	return l.MaxSize
}

// GetMaxBackups 最大日志文件数
func (l *LogConfig) GetMaxBackups() int {
	if l == nil || l.MaxBackups == 0 {
		return 10
	}
	return l.MaxBackups
}

// GetCompress 文件是否压缩
func (l *LogConfig) GetCompress() bool {
	if l == nil {
		return false
	}
	return l.Compress
}
