package elastic

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/wpliam/common-wrap/client"
)

func init() {
	client.Register("elastic", &Client{})
}

type Client struct {
}

func (c *Client) Invoke(ctx context.Context, opts ...client.Option) error {
	o := &client.Options{}
	for _, opt := range opts {
		opt(o)
	}
	cli, err := elastic.NewClient(
		elastic.SetURL(o.Target),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(o.Username, o.Password),
	)
	if err != nil {
		return err
	}
	_, _, err = cli.Ping(o.Target).Do(ctx)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	connPool[o.Name] = cli
	return nil
}
