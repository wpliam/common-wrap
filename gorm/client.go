package gorm

import (
	"context"
	"github.com/wpliap/common-wrap/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	client.Register("gorm", &Client{})
}

type Client struct {
}

func (c *Client) Invoke(ctx context.Context, opts ...client.Option) error {
	o := &client.Options{}
	for _, opt := range opts {
		opt(o)
	}
	db, err := gorm.Open(mysql.Open(o.Target), &gorm.Config{})
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	connPool[o.Name] = db
	return nil
}
