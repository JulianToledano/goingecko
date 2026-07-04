package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// CoinsOhlcRange allows you to get the OHLC chart (Open, High, Low, Close) of a coin within a range of timestamps.
//
// This is a Pro API endpoint.
func (c *CoinsClient) CoinsOhlcRange(ctx context.Context, id, vsCurrency, from, to, interval string) (*types.Ohlc, error) {
	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", from)
	params.Add("to", to)
	params.Add("interval", interval)

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "ohlc/range", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.Ohlc
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
