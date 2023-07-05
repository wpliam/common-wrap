package log

type LoggerConfig struct {
	CallerSkip  int         `yaml:"caller_skip"`
	WriteConfig WriteConfig `yaml:"write_config"`
}

type WriteConfig struct {
	Filename   string `yaml:"filename"`
	LogPath    string `yaml:"log_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}
