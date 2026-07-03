package api

import (
	"net/url"
	"testing"
)

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
