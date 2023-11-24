package test

import (
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestNftsList(t *testing.T) {
	cgClient := goingecko.NewClient(nil)
	data, err := cgClient.NftsList("", "", 0, 0)
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestNftsId(t *testing.T) {
	cgClient := goingecko.NewClient(nil)
	data, err := cgClient.NftsId("squiggly")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestNftsContract(t *testing.T) {
	cgClient := goingecko.NewClient(nil)
	data, err := cgClient.NftsContract("ethereum", "0x36F379400DE6c6BCDF4408B282F8b685c56adc60")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
