package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JulianToledano/goingecko/v3/api/internal"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// idHistoryOption is an interface that extends internal.Option to provide
// specific options for the CoinsIdHistory endpoint. It includes a marker
// method isIdHistoryOption() to ensure type safety for history-specific options.
type idHistoryOption interface {
	internal.Option
	isIdHistoryOption()
}

// WithLocalizationIdHistoryOption sets whether to include localized data.
// If true, returns localized data in response (name, description, etc.)
// If false, returns data in English.
func WithLocalizationIdHistoryOption(loc bool) idHistoryOption {
	return localizationIdHistoryOption{loc}
}

// CoinsIdHistory allows you to query the historical data (price, market cap, 24hrs volume, etc) at a given date for a
// coin based on a particular coin id.
//
// 👍 Tips
//
//	You may obtain the coin id (api id) via several ways:
//	refers to respective coin page and find ‘api id’
//	refers to /coins/list endpoint
//	refers to google sheets here
//
// 📘 Notes
//
//	The data returned is at 00:00:00 UTC
//	The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35)
func (c *CoinsClient) CoinsIdHistory(ctx context.Context, id, date string, options ...idHistoryOption) (*types.History, error) {
	params := url.Values{}
	params.Set("date", date)

	// Apply all the options
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s/%s?%s", c.coinsUrl(), id, "history", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.History
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type localizationIdHistoryOption struct{ localization bool }

func (o localizationIdHistoryOption) Apply(v *url.Values) {
	v.Set("localization", strconv.FormatBool(o.localization))
}
func (localizationIdHistoryOption) isIdHistoryOption() {}
