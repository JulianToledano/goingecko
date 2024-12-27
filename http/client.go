package http

import (
	"context"
	"io"
	"net/http"
)

type (
	ApiHeaderFn func(r *http.Request)

	option func(*Client) *Client
)

func WithHttpClient(client *http.Client) option {
	return func(c *Client) *Client {
		c.httpClient = client
		return c
	}
}

func WithApiHeaderFn(fn ApiHeaderFn) option {
	return func(c *Client) *Client {
		c.apiHeaderSetter = fn
		return c
	}
}

type Client struct {
	httpClient      *http.Client
	apiHeaderSetter ApiHeaderFn
}

func NewClient(opts ...option) *Client {
	c := &Client{}

	for _, opt := range opts {
		c = opt(c)
	}

	return c
}

func (c *Client) MakeReq(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.apiHeaderSetter != nil {
		c.apiHeaderSetter(req)
	}

	_, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) do(req *http.Request) (int, []byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}

	return resp.StatusCode, body, nil
}
