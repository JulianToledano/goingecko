package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type (
	// ApiHeaderFn is a function type that sets headers on HTTP requests
	ApiHeaderFn func(*http.Request)

	// option is a function type that configures a Client instance
	option func(*Client) *Client
)

// WithHttpClient returns an option that sets the HTTP client used by the Client
func WithHttpClient(client *http.Client) option {
	return func(c *Client) *Client {
		c.httpClient = client
		return c
	}
}

// WithApiHeaderFn returns an option that sets a function to add API headers to requests
func WithApiHeaderFn(fn ApiHeaderFn) option {
	return func(c *Client) *Client {
		c.apiHeaderSetter = fn
		return c
	}
}

// Client is an HTTP client that can make API requests with optional headers
type Client struct {
	// httpClient is the underlying HTTP client used to make requests
	httpClient *http.Client
	// apiHeaderSetter is an optional function to set API headers on requests
	apiHeaderSetter ApiHeaderFn
}

// NewClient creates a new Client with the given options
func NewClient(opts ...option) *Client {
	c := &Client{}

	for _, opt := range opts {
		c = opt(c)
	}

	return c
}

// APIError is an error returned by the client when the API returns a non-200 status code
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api returned status code %d", e.StatusCode)
}

// MakeReq makes a GET request to the given URL with the configured client
// It returns the response body as bytes or an error if the request fails
func (c *Client) MakeReq(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.apiHeaderSetter != nil {
		c.apiHeaderSetter(req)
	}

	code, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, &APIError{
			StatusCode: code,
			Body:       resp,
		}
	}

	return resp, nil
}

// do executes an HTTP request and returns the status code, response body and any error
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
