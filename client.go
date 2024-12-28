package goingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JulianToledano/goingecko/ping"
)

const (
	apiHeader    = "x-cg-demo-api-key"
	proApiHeader = "x-cg-pro-api-key"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
	apiKey     string
	apiHeader  string
}

func NewClient(httpClient *http.Client, apiKey string, isPro ...bool) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if isPro != nil {
		return &Client{
			httpClient: httpClient,
			baseUrl:    ProBaseURL,
			apiKey:     apiKey,
			apiHeader:  proApiHeader,
		}
	}

	return &Client{
		httpClient: httpClient,
		baseUrl:    BaseURL,
		apiKey:     apiKey,
		apiHeader:  apiHeader,
	}
}

func (c *Client) Close() {
	c.httpClient.CloseIdleConnections()
}

func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// MakeReq HTTP request helper
func (c *Client) MakeReq(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if c.apiKey != "" {
		req.Header.Add(c.apiHeader, c.apiKey)
	}

	if err != nil {
		return nil, err
	}
	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Ping /ping endpoint
func (c *Client) Ping(ctx context.Context) (*ping.Ping, error) {
	resp, err := c.MakeReq(ctx, fmt.Sprintf("%s/ping", c.baseUrl))
	if err != nil {
		return nil, err
	}
	var data *ping.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
