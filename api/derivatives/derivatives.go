package derivatives

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/api/derivatives/types"
)

// Derivatives allows you to query all the tickers from derivatives exchanges on CoinGecko.
//
//	📘 Notes
//
//	    Data for open_interest and volume_24h in the endpoint responses are in USD
//	    Cache / Update Frequency: every 30 seconds for all the API plans
func (c *DerivativesClient) Derivatives(ctx context.Context) ([]types.Derivative, error) {

	rUrl := fmt.Sprintf("%s", c.derivativesUrl())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.Derivative
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
