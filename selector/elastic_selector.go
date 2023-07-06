package selector

import (
	"context"
	"github.com/wpliam/common-wrap/registry"

	"github.com/olivere/elastic/v7"
)

func init() {
	Register("elastic", &elasticSelector{})
}

type elasticSelector struct {
}

func (e *elasticSelector) Select(opt ...Option) (registry.Proxy, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}
	client, err := elastic.NewClient(
		elastic.SetURL(opts.Target),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(opts.Username, opts.Password),
	)
	if err != nil {
		return nil, err
	}
	_, _, err = client.Ping(opts.Target).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return registry.NewElasticProxy(client), nil
}
