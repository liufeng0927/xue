package mini

import "context"

type Config struct {
	AppId     string
	AppSecret string
	BaseUri   string
}

type Client struct {
	ctx       context.Context
	AppId     string
	AppSecret string
	BaseUri   string
}

func NewClient(ctx context.Context, c *Config) *Client {
	return &Client{
		ctx:       ctx,
		AppId:     c.AppId,
		AppSecret: c.AppSecret,
		BaseUri:   c.BaseUri,
	}
}
