package jam

import (
	"context"
	"fmt"

	"github.com/wpliam/common-wrap/client"
	"github.com/wpliam/common-wrap/config"
	"github.com/wpliam/common-wrap/server"

	_ "github.com/wpliam/common-wrap/elastic"
	_ "github.com/wpliam/common-wrap/gorm"
	_ "github.com/wpliam/common-wrap/log"
	_ "github.com/wpliam/common-wrap/redis"
)

func New() *server.Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("load config err " + err.Error())
	}
	if cfg.Plugins != nil {
		if err = cfg.Plugins.Setup(); err != nil {
			panic("plugin setup err" + err.Error())
		}
	}

	for _, s := range cfg.Client.Service {
		cli := client.Get(s.Protocol)
		opts := []client.Option{
			client.WithName(s.Name),
			client.WithTarget(s.Target),
			client.WithUsername(s.Username),
			client.WithPassword(s.Password),
			client.WithProtocol(s.Protocol),
		}
		if err = cli.Invoke(context.Background(), opts...); err != nil {
			panic(fmt.Sprintf("client [%s] invoke err:%v", s.Name, err))
		}
	}
	svr := &server.Server{}
	for _, s := range cfg.Server.Service {
		svr.AddService(s.Name, server.NewService(s.Name, server.WithPort(s.Port), server.WithProtocol(s.Protocol)))
	}
	return svr
}
