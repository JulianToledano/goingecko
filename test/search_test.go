package test

import (
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestSearch(t *testing.T) {
	cgClient := goingecko.NewClient(nil)
	data, err := cgClient.Search("bitcoin")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
