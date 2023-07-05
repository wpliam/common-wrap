package selector

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func init() {
	Register("elastic", &elasticSelector{})
}

type elasticSelector struct {
}

func (e *elasticSelector) Select(opt ...Option) (interface{}, error) {
	opts := &Options{}
	for _, o := range opt {
		o(opts)
	}
	proxy, err := elastic.NewClient(
		elastic.SetURL(opts.Target),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(opts.Username, opts.Password),
	)
	if err != nil {
		return nil, err
	}
	_, _, err = proxy.Ping(opts.Target).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return proxy, nil
}
