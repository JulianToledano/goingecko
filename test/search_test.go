package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko/v2"
)

func TestSearch(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.Search(context.Background(), "bitcoin")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
