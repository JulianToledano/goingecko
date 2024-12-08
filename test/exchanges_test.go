package test

import (
	"context"
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestExchanges(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	data, err := cgClient.Exchanges(context.Background(), "", "")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesList(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesList(context.Background())
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesId(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesId(context.Background(), "sushiswap")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesIdTickers(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesIdTickers(context.Background(), "sushiswap", "", "", 0, "", "")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestExchangesIdVolumeChart(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.ExchangesIdVolumeChart(context.Background(), "sushiswap", "1")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
