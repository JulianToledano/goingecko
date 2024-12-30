package internal

import geckohttp "github.com/JulianToledano/goingecko/http"

type Client struct {
	*geckohttp.Client

	URL string
}

func NewClient(httpClient *geckohttp.Client, url string) *Client {
	return &Client{
		Client: httpClient,
		URL:    url,
	}
}
