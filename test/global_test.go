package test

import (
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestGlobal(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	r, err := cgClient.Global()
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDefi(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	r, err := cgClient.DecentrilizedFinanceDEFI()
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
