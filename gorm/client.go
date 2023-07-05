package gorm

import (
	"github.com/wpliam/common-wrap/client"
	"gorm.io/gorm"
)

var NewClientProxy = func(name string, opts ...client.Option) *gorm.DB {
	options := make([]client.Option, 0, len(opts)+1)
	options = append(options, opts...)
	options = append(options,
		client.WithProtocol("gorm"),
	)
	newClient := client.New()
	if err := newClient.Invoke(name, options...); err != nil {
		panic("gorm client invoke err" + err.Error())
	}
	proxy := newClient.Get()
	if proxy == nil {
		panic("get gorm client proxy not exist")
	}
	return proxy.(*gorm.DB)
}
