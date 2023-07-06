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
	proxy := client.New()
	if err := proxy.Invoke(name, options...); err != nil {
		panic("elastic client invoke err" + err.Error())
	}
	return proxy.Get().Elastic()
}
