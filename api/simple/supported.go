package simple

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/api/simple/types"
)

// SimpleSupportedVsCurrency allows you to query all the supported currencies on CoinGecko.
//
//	👍 Tips
//
//	    You may use this endpoint to query the list of currencies for other endpoints that contain params like vs_currencies
//
//	📘Notes
//
//	    Cache/Update Frequency: every 30 seconds for Pro API (Analyst, Lite, Pro, Enterprise)
func (c *SimpleClient) SimpleSupportedVsCurrency(ctx context.Context) (*types.SupportedVsCurrency, error) {
	rUrl := fmt.Sprintf("%s/%s", c.simpleUrl(), "supported_vs_currencies")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.SupportedVsCurrency
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
