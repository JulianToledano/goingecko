package api

import (
	"context"
	"net/url"
	"testing"
)

type dictOptionTestClient struct {
	requestURL string
	response   []byte
}

func (c *dictOptionTestClient) MakeReq(_ context.Context, requestURL string) ([]byte, error) {
	c.requestURL = requestURL
	return c.response, nil
}

func TestWithURLParams(t *testing.T) {
	params := url.Values{}

	WithURLParams(map[string]any{
		"page":      2,
		"sparkline": true,
		"precision": "full",
	}).Apply(&params)

	if got := params.Get("page"); got != "2" {
		t.Fatalf("page = %q, want %q", got, "2")
	}
	if got := params.Get("sparkline"); got != "true" {
		t.Fatalf("sparkline = %q, want %q", got, "true")
	}
	if got := params.Get("precision"); got != "full" {
		t.Fatalf("precision = %q, want %q", got, "full")
	}
}

func TestWithURLParamsAppliesToRealRequest(t *testing.T) {
	httpClient := &dictOptionTestClient{
		response: []byte(`[]`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	_, err := c.CoinsMarket(context.Background(), "usd", WithURLParams(map[string]any{
		"per_page":  25,
		"threshold": 1.5,
		"sparkline": true,
		"order":     "market_cap_desc",
	}))
	if err != nil {
		t.Fatalf("CoinsMarket() error = %v", err)
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/coins/markets" {
		t.Fatalf("request path = %q, want /api/v3/coins/markets", requestURL.Path)
	}

	query := requestURL.Query()
	if got := query.Get("vs_currency"); got != "usd" {
		t.Errorf("vs_currency = %q, want usd", got)
	}
	if got := query.Get("per_page"); got != "25" {
		t.Errorf("per_page = %q, want 25", got)
	}
	if got := query.Get("threshold"); got != "1.5" {
		t.Errorf("threshold = %q, want 1.5", got)
	}
	if got := query.Get("sparkline"); got != "true" {
		t.Errorf("sparkline = %q, want true", got)
	}
	if got := query.Get("order"); got != "market_cap_desc" {
		t.Errorf("order = %q, want market_cap_desc", got)
	}
}
