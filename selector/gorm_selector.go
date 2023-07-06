package selector

import (
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"github.com/wpliam/common-wrap/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	Register("gorm", &gormSelector{})
}

type gormSelector struct {
}

type Proxy interface {
	*gorm.DB | *redis.Client | *elastic.Client
}

func (g *gormSelector) Select(opt ...Option) (registry.Proxy, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}
	db, err := gorm.Open(mysql.Open(opts.Target), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return registry.NewGormProxy(db), nil
}
