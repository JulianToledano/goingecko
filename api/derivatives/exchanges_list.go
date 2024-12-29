package derivatives

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/api/derivatives/types"
)

// DerivativesExchangesList allows you to query all the derivatives exchanges with id and name on CoinGecko.
//
//	👍 Tips
//
//	    You may use this endpoint to query the list of exchanges for other endpoints that contain params like id (derivatives exchange's id)
//
//	📘 Notes
//
//	    Cache / Update Frequency: every 5 minutes for all the API plans
func (c *DerivativesClient) DerivativesExchangesList(ctx context.Context) ([]types.DerivativesListItem, error) {
	rUrl := fmt.Sprintf("%s/exchanges/list", c.derivativesUrl())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.DerivativesListItem
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
