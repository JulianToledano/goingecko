package coins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/api"

	cointypes "github.com/JulianToledano/goingecko/api/coins/types"
)

// listOption is an interface that extends api.Option to provide
// specific options for the CoinsList endpoint. It includes a marker
// method isListOption() to ensure type safety for list-specific options.
type listOption interface {
	api.Option
	isListOption()
}

// WithIncludePlatform include platform and token's contract addresses, default: false.
func WithIncludePlatform(include bool) listOption { return includePlatformOption{include: include} }

// WithStatus filter by status of coins, default: active
// valid values: active and inactive.
func WithStatus(status string) listOption { return statusOption{status: status} }

// CoinsList allows you to query all the supported coins on CoinGecko with coins id, name and symbol.
//
// 👍 Tips
// You may use this endpoint to query the list of coins with coin id for other endpoints that contain params like id or
// ids (coin id).
// By default, this endpoint returns full list of active coins that are currently listed on CoinGecko.com , you can also
// flag status=inactive to retrieve coins that are no longer available on CoinGecko.com . The inactive coin ids can also
// be used with selected historical data endpoints.
//
// 📘 Notes
// There is no pagination required for this endpoint
// Cache/Update Frequency: Every 5 minutes for all the API plans
func (c *Client) CoinsList(ctx context.Context, options ...listOption) ([]*cointypes.CoinInfo, error) {
	params := url.Values{}

	// Apply all the options
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/%s?%s", c.coinsUrl(), "list", params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []*cointypes.CoinInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type includePlatformOption struct{ include bool }

func (o includePlatformOption) isListOption() {}
func (o includePlatformOption) Apply(v *url.Values) {
	v.Set("include_platform", strconv.FormatBool(o.include))
}

type statusOption struct{ status string }

func (o statusOption) isListOption() {}
func (o statusOption) Apply(v *url.Values) {
	v.Set("status", o.status)
}
