package exchanges

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/exchanges/types"
)

// ExchangesIdVolumeChart allows you to query the historical volume chart data with time in UNIX and trading volume data in BTC based on exchange’s id.
//
//	📘 Notes
//
//	    You can use this endpoint to query the historical volume chart data of derivatives exchanges as well
//	    The exchange volume in the response is provided in BTC. To convert it to other currencies, please use /exchange_rates endpoint
//	    Data granularity is automatic (cannot be adjusted)
//	        1 day = 10-minutely
//	        7, 14 days = hourly
//	        30 days & above = daily
//	    Cache / Update Frequency: every 60 seconds for all the API plans
func (c *ExchangesClient) ExchangesIdVolumeChart(ctx context.Context, id, days string) ([]types.Volume, error) {
	if _, ok := types.VolumeChartValidDays[days]; !ok {
		return nil, errors.New("not valid day")
	}

	params := url.Values{}
	params.Add("days", days)

	rUrl := fmt.Sprintf("%s/%s/volume_chart?%s", c.exchangesUrl(), id, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.Volume
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
