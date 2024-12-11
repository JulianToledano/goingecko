package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestContractInfo(t *testing.T) {
	coin := "ethereum"
	contract := "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"

	cgClient := goingecko.NewClient(nil, "")

	contractData, err := cgClient.ContractInfo(context.Background(), coin, contract)
	if err != nil {
		t.Errorf("Error: %v", err)
	} else if contractData.ID != "basic-attention-token" {
		t.Errorf("Id: %s", contractData.ID)
	}

}

func TestContractMarketChart(t *testing.T) {
	coin := "ethereum"
	contract := "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"

	cgClient := goingecko.NewClient(nil, "")

	contractData, err := cgClient.ContractMarketChart(context.Background(), coin, contract, "usd", "10")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if contractData == nil {
		t.Errorf("Error nil")
	}
}

func TestContractMarketChartRange(t *testing.T) {
	coin := "ethereum"
	contract := "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"

	cgClient := goingecko.NewClient(nil, "")

	contractData, err := cgClient.ContractMarketChartRange(context.Background(), coin, contract, "usd", "1392500000", "1422577232")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if contractData == nil {
		t.Errorf("Error nil")
	}
}
