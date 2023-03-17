package ftp

import (
	"fmt"
	"github.com/wpliap/common-wrap/log"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/wpliap/common-wrap/config"
)

// NewFtpProxy 创建一个ftp代理
func NewFtpProxy(name string, opt ...config.Option) *ftp.ServerConn {
	cfg := config.GetConnConf(name, opt...)
	addr := fmt.Sprintf("%s:%d", cfg.GetHost(), cfg.GetPort())
	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(time.Duration(cfg.GetTimeout())*time.Millisecond))
	if err != nil {
		panic("ftp dial err " + err.Error())
	}
	if err = conn.Login(cfg.GetUsername(), cfg.GetPassword()); err != nil {
		panic("ftp login err " + err.Error())
	}
	log.Infof("create ftp proxy success name:%s host:%s port:%d", name, cfg.GetHost(), cfg.GetPort())
	return conn
}
