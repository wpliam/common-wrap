package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/wpliap/common-wrap/config"
	"github.com/wpliap/common-wrap/log"
)

// NewElasticProxy 创建es代理
func NewElasticProxy(name string, opt ...config.Option) *elastic.Client {
	cfg := config.GetConnConf(name, opt...)
	url := fmt.Sprintf("http://%s:%d", cfg.GetHost(), cfg.GetPort())
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(url),
	)
	if err != nil {
		panic("elastic new client err " + err.Error())
	}
	_, _, err = client.Ping(url).Do(context.Background())
	if err != nil {
		panic("elastic ping err " + err.Error())
	}
	log.Infof("create elastic proxy success name:%s host:%s port:%d", name, cfg.GetHost(), cfg.GetPort())
	return client
}
