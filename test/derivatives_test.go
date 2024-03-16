package test

import (
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestDerivatives(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.Derivatives()
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDerivativesExchanges(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.DerivativesExchanges("", 0, 0)
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDerivativesExchangesId(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.DerivativesExchangesId("binance_futures", "all")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDerivativesExchangesList(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.DerivativesExchangesList()
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
