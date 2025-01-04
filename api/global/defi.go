package global

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/global/types"
)

// DecentralizedFinanceDEFI allows you query top 100 cryptocurrency global decentralized finance (defi) data including defi market cap, trading volume.
//
//	📘Notes
//
//	    Cache / Update Frequency: every 60 minutes for all the API plans
func (c *GlobalClient) DecentralizedFinanceDEFI(ctx context.Context) (*types.Defi, error) {
	rUrl := fmt.Sprintf("%s/decentralized_finance_defi", c.globalUrl())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Defi
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
