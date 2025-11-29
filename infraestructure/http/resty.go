package http

import "github.com/go-resty/resty/v2"

type Client struct {
	client *resty.Client
	URL string
}

type Config struct {
	Timeout string
	EnableTraces bool
}

func NewClient(url string, config Config) *Client {
	return &Client{
		client: resty.New(),
		URL: url,
	}
}

func (c *Client) Do() (*resty.Response, error) {
	return c.client.R().Get(c.URL)
}