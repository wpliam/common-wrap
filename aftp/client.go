package aftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/wpliap/common-wrap/config"
	"time"
)

// NewFtpProxy 创建一个ftp代理
func NewFtpProxy(name string) *ftp.ServerConn {
	cfg := config.GetConnConf(name)
	if cfg == nil {
		panic(name + " ftp conf not exist")
	}
	fmt.Println("timeout:", cfg.GetTimeout())
	addr := fmt.Sprintf("%s:%d", cfg.GetHost(), cfg.GetPort())
	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(time.Duration(cfg.GetTimeout())*time.Millisecond))
	if err != nil {
		panic("ftp dial err " + err.Error())
	}
	if err = conn.Login(cfg.GetUsername(), cfg.GetPassword()); err != nil {
		panic("ftp login err " + err.Error())
	}
	return conn
}
