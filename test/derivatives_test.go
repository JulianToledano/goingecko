package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestDerivatives(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.Derivatives(context.Background())
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDerivativesExchanges(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.DerivativesExchanges(context.Background(), "", 0, 0)
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDerivativesExchangesId(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.DerivativesExchangesId(context.Background(), "binance_futures", "all")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDerivativesExchangesList(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.DerivativesExchangesList(context.Background())
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
