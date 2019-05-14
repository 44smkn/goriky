package client

import (
	"context"
	"net/http"
	"net/url"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	// logger
}

func New() *Client {

	httpClient := &http.Client{}
	client := &Client{
		HTTPClient: httpClient,
	}
	return client
}

func (c *Client) SendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
