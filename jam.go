package jam

import (
	"github.com/wpliam/common-wrap/client"
	"github.com/wpliam/common-wrap/config"
	"github.com/wpliam/common-wrap/server"

	_ "github.com/wpliam/common-wrap/log"
	_ "github.com/wpliam/common-wrap/selector"
)

func New() *server.Server {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("load config err " + err.Error())
	}
	// 设置插件
	if err = cfg.Plugins.Setup(); err != nil {
		panic("plugin setup err" + err.Error())
	}
	for _, s := range cfg.Client.Service {
		client.RegisterClientConfig(s.Name, &s)
	}
	svr := &server.Server{}
	for _, s := range cfg.Server.Service {
		opts := []server.Option{
			server.WithPort(s.Port),
			server.WithProtocol(s.Protocol),
		}
		svr.AddService(s.Name, server.New(s.Name, opts...))
	}
	return svr
}
