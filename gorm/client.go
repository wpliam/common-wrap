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
	proxy := client.New()
	if err := proxy.Invoke(name, options...); err != nil {
		panic("gorm client invoke err" + err.Error())
	}
	return proxy.Get().Gorm()
}
