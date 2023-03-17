package gorm

import (
	"fmt"
	"github.com/wpliap/common-wrap/log"

	"github.com/wpliap/common-wrap/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormProxy 创建Gorm代理
func NewGormProxy(name string, opt ...config.Option) *gorm.DB {
	cfg := config.GetConnConf(name, opt...)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.GetUsername(), cfg.GetPassword(), cfg.GetHost(), cfg.GetPort(), cfg.GetDatabase())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gorm open err " + err.Error())
	}
	log.Infof("create gorm proxy success name:%s host:%s port:%d", name, cfg.GetHost(), cfg.GetPort())
	return db
}
