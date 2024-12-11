package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestExchangeRates(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	r, err := cgClient.ExchangeRates(context.Background())
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
