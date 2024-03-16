package test

import (
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestExchanges(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	data, err := cgClient.Exchanges("", "")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesList(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesList()
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesId(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesId("sushiswap")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesIdTickers(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesIdTickers("sushiswap", "", "", 0, "", "")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesIdVolumeChart(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesIdVolumeChart("sushiswap", "1")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
