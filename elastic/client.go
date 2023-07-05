package elastic

import (
	"github.com/olivere/elastic/v7"
	"github.com/wpliam/common-wrap/client"
)

var NewClientProxy = func(name string, opts ...client.Option) *elastic.Client {
	options := make([]client.Option, 0, len(opts)+1)
	options = append(options, opts...)
	options = append(options,
		client.WithProtocol("elastic"),
	)
	newClient := client.New()
	if err := newClient.Invoke(name, options...); err != nil {
		panic("elastic client invoke err" + err.Error())
	}
	proxy := newClient.Get()
	if proxy == nil {
		panic("get elastic client proxy not exist")
	}
	return proxy.(*elastic.Client)
}
