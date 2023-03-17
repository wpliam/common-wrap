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
	Timeout  int    `yaml:"timeout"`
}

type Option func(*ConnConf)

// GetConnConf 获取连接配置
func GetConnConf(name string, opt ...Option) *ConnConf {
	connConf := &ConnConf{}
	for _, conn := range GetConf().Client.Service {
		if name == conn.Name {
			connConf = conn
		}
	}
	for _, o := range opt {
		o(connConf)
	}
	return connConf
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

// GetTimeout 获取超时
func (c *ConnConf) GetTimeout() int {
	if c == nil {
		return 2000
	}
	return c.Timeout
}

func WithHost(host string) Option {
	return func(o *ConnConf) {
		o.Host = host
	}
}

func WithPort(port uint16) Option {
	return func(o *ConnConf) {
		o.Port = port
	}
}

func WithUsername(username string) Option {
	return func(o *ConnConf) {
		o.Username = username
	}
}

func WithPassword(password string) Option {
	return func(o *ConnConf) {
		o.Password = password
	}
}

func WithDatabase(database string) Option {
	return func(o *ConnConf) {
		o.Database = database
	}
}

func WithTimeout(timeout int) Option {
	return func(o *ConnConf) {
		o.Timeout = timeout
	}
}
