package api

import (
	"context"
	"testing"
)

func TestClient_CoinsList(t *testing.T) {
	c := NewDefaultClient()

	got, err := c.CoinsList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("nil response")
	}
}
