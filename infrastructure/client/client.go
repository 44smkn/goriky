package client

import (
	"context"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	// logger
}

func New(urlStr string) (*Client, error) {

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	httpClient := &http.Client{}

	client := &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}
	return client, nil
}

func (c *Client) SendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
