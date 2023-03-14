package agorm

import (
	"fmt"

	"github.com/wpliap/common-wrap/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormProxy 创建Gorm代理
func NewGormProxy(name string) *gorm.DB {
	cfg := config.GetConnConf(name)
	if cfg == nil {
		panic(name + " mysql conf not exist")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.GetUsername(), cfg.GetPassword(), cfg.GetHost(), cfg.GetPort(), cfg.GetDatabase())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gorm open err " + err.Error())
	}
	return db
}
