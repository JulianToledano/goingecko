package test

import (
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestTrending(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	r, err := cgClient.Trending()
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
