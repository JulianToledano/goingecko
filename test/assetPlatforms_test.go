package test

import (
	"github.com/JulianToledano/goingecko"
	"testing"
)

func TestAssetPlatforms(t *testing.T) {

	cgClient := goingecko.NewClient(nil, "")

	assetData, err := cgClient.AssetPlatforms("")
	if assetData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
