package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko/v2"
)

func TestTrending(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	r, err := cgClient.Trending(context.Background())
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
