package registry

import (
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Proxy interface {
	Gorm() *gorm.DB
	Redis() *redis.Client
	Elastic() *elastic.Client
}

func NewGormProxy(db *gorm.DB) Proxy {
	return &proxy{gdb: db}
}

func NewRedisProxy(client *redis.Client) Proxy {
	return &proxy{rdb: client}
}

func NewElasticProxy(client *elastic.Client) Proxy {
	return &proxy{es: client}
}

type proxy struct {
	gdb *gorm.DB
	rdb *redis.Client
	es  *elastic.Client
}

func (p *proxy) Redis() *redis.Client {
	return p.rdb
}

func (p *proxy) Elastic() *elastic.Client {
	return p.es
}

func (p *proxy) Gorm() *gorm.DB {
	return p.gdb
}
