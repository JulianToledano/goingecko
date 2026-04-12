package websocket

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	coderws "github.com/coder/websocket"
)

const (
	// BaseURL is CoinGecko's websocket endpoint.
	BaseURL = "wss://stream.coingecko.com/v1"

	proAPIKeyHeaderKey = "x-cg-pro-api-key"
	proAPIKeyQueryKey  = "x_cg_pro_api_key"
)

var (
	// ErrMissingAPIKey is returned when the websocket client is missing a Pro API key.
	ErrMissingAPIKey = errors.New("coingecko websocket api key is required")
	// ErrInvalidAuthMode is returned when the websocket client is configured with an unsupported auth mode.
	ErrInvalidAuthMode = errors.New("invalid websocket auth mode")
)

// AuthMode controls how the CoinGecko Pro API key is sent during the websocket handshake.
type AuthMode string

const (
	// AuthModeQuery appends the Pro API key as the x_cg_pro_api_key query parameter.
	AuthModeQuery AuthMode = "query"
	// AuthModeHeader sends the Pro API key in the x-cg-pro-api-key header.
	AuthModeHeader AuthMode = "header"
)

// Option configures a websocket client.
type Option func(*Client)

// Client manages websocket connection setup for CoinGecko's streaming API.
type Client struct {
	apiKey     string
	url        string
	authMode   AuthMode
	header     http.Header
	httpClient *http.Client
}

// NewClient creates a websocket client for CoinGecko's streaming API.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:   strings.TrimSpace(apiKey),
		url:      BaseURL,
		authMode: AuthModeQuery,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithURL configures a custom websocket endpoint.
func WithURL(rawURL string) Option {
	return func(c *Client) {
		if rawURL == "" {
			return
		}
		c.url = rawURL
	}
}

// WithAuthMode configures whether authentication is sent as a query parameter or request header.
func WithAuthMode(mode AuthMode) Option {
	return func(c *Client) {
		c.authMode = mode
	}
}

// WithHeader merges additional headers into the websocket handshake request.
func WithHeader(header http.Header) Option {
	return func(c *Client) {
		if len(header) == 0 {
			return
		}
		if c.header == nil {
			c.header = make(http.Header)
		}
		for key, values := range header {
			for _, value := range values {
				c.header.Add(key, value)
			}
		}
	}
}

// WithHTTPClient configures the HTTP client used during the websocket handshake.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// Connect establishes a websocket connection to CoinGecko's streaming API.
func (c *Client) Connect(ctx context.Context) (*Conn, error) {
	if strings.TrimSpace(c.apiKey) == "" {
		return nil, ErrMissingAPIKey
	}

	dialURL, header, err := c.dialConfig()
	if err != nil {
		return nil, err
	}

	opts := &coderws.DialOptions{
		HTTPClient: c.httpClient,
		HTTPHeader: header,
	}

	conn, _, err := coderws.Dial(ctx, dialURL, opts)
	if err != nil {
		return nil, err
	}

	return &Conn{conn: conn}, nil
}

func (c *Client) dialConfig() (string, http.Header, error) {
	header := cloneHeader(c.header)

	switch c.authMode {
	case AuthModeQuery:
		u, err := url.Parse(c.url)
		if err != nil {
			return "", nil, err
		}
		query := u.Query()
		query.Set(proAPIKeyQueryKey, c.apiKey)
		u.RawQuery = query.Encode()
		return u.String(), header, nil
	case AuthModeHeader:
		header.Set(proAPIKeyHeaderKey, c.apiKey)
		return c.url, header, nil
	default:
		return "", nil, ErrInvalidAuthMode
	}
}

func cloneHeader(header http.Header) http.Header {
	if header == nil {
		return make(http.Header)
	}
	return header.Clone()
}

// FormatSubscriptionTarget returns the network-scoped identifier CoinGecko expects for on-chain subscriptions.
func FormatSubscriptionTarget(networkID, address string) string {
	return strings.TrimSpace(networkID) + ":" + strings.TrimSpace(address)
}
