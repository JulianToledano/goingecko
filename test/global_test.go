package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko/v2"
)

func TestGlobal(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	r, err := cgClient.Global(context.Background())
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDefi(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	r, err := cgClient.DecentralizedFinanceDEFI(context.Background())
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
