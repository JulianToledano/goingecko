package http

import (
	"context"
	"io"
	"net/http"
	"time"
)

// ApiHeaderFn is a function type that sets headers on HTTP requests
type ApiHeaderFn func(*http.Request)

// HttpClient defines the interface for making HTTP requests.
// It provides a method to make GET requests and return the response body as bytes.
type HttpClient interface {
	// MakeReq makes a GET request to the given URL.
	// It returns the response body as bytes and any error that occurred.
	// The context can be used to cancel the request or set a timeout.
	MakeReq(ctx context.Context, url string) ([]byte, error)
}

var _ HttpClient = &Client{}

// Client is an HTTP client that can make API requests with optional headers and rate limiting
type Client struct {
	// httpClient is the underlying HTTP client used to make requests
	httpClient *http.Client
	// apiHeaderSetter is an optional function to set API headers on requests
	apiHeaderSetter ApiHeaderFn
}

// NewClient creates a new Client with the given options
func NewClient(opts ...option[*Client]) *Client {
	c := &Client{}

	for _, opt := range opts {
		c = opt(c)
	}

	// Set default HTTP client if none provided
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return c
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
