package test

import (
	"context"
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestPublicTreasuryCoinId(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.PublicTreasuryCoinId(context.Background(), "bitcoin")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
