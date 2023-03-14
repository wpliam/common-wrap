package config

// ClientConfig 客户端配置
type ClientConfig struct {
	Service []*ConnConf `yaml:"service"`
}

// ConnConf 连接配置
type ConnConf struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// GetConnConf 获取连接配置
func GetConnConf(name string) *ConnConf {
	for _, conn := range GetConf().Client.Service {
		if name == conn.Name {
			return conn
		}
	}
	return nil
}

// GetName 获取名称
func (c *ConnConf) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

// GetHost 获取host
func (c *ConnConf) GetHost() string {
	if c == nil {
		return ""
	}
	return c.Host
}

// GetPort 获取端口
func (c *ConnConf) GetPort() uint16 {
	if c == nil {
		return 0
	}
	return c.Port
}

// GetUsername 获取用户名
func (c *ConnConf) GetUsername() string {
	if c == nil {
		return ""
	}
	return c.Username
}

// GetPassword 获取密码
func (c *ConnConf) GetPassword() string {
	if c == nil {
		return ""
	}
	return c.Password
}

// GetDatabase 获取db
func (c *ConnConf) GetDatabase() string {
	if c == nil {
		return ""
	}
	return c.Database
}
