package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestSimplePrice(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")

	r, _ := cgClient.SimplePrice(context.Background(), "bitcoin,ethereum", "btc,eth", true, true, true, true)
	if r == nil {
		t.Errorf("Error")
	}

}

func TestSimpleTokenPrice(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")

	r, _ := cgClient.SimpleTokenPrice(context.Background(), "ethereum", "0x0D8775F648430679A709E98d2b0Cb6250d2887EF", "btc,eth", true, true, true, true)
	if r == nil {
		t.Errorf("Error")
	}

}

func TestSimpleSupportedVsCurrency(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")

	supVcurr, _ := cgClient.SimpleSupportedVsCurrency(context.Background())
	if supVcurr == nil {
		t.Errorf("Error")
	}
}
