package test

import (
	"context"
	"testing"

	"github.com/JulianToledano/goingecko/v2"
)

func TestCategoriesList(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	categoriesList, err := cgClient.CategoriesList(context.Background())
	if categoriesList == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCategories(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	categoriesList, err := cgClient.Categories(context.Background(), "market_cap_desc")
	if categoriesList == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
