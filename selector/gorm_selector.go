package selector

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	Register("gorm", &gormSelector{})
}

type gormSelector struct {
}

func (g *gormSelector) Select(opt ...Option) (interface{}, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}
	db, err := gorm.Open(mysql.Open(opts.Target), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
