package global

import (
	"context"
	"encoding/json"

	"github.com/JulianToledano/goingecko/v3/api/global/types"
)

// Global allows you query cryptocurrency global data including active cryptocurrencies, markets, total crypto market cap and etc.
//
//	📘Notes
//
//	    Cache / Update Frequency: every 10 minutes for all the API plans
func (c *GlobalClient) Global(ctx context.Context) (*types.Global, error) {
	resp, err := c.MakeReq(ctx, c.globalUrl())
	if err != nil {
		return nil, err
	}

	var data *types.Global
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
