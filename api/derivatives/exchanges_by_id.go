package derivatives

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/derivatives/types"
	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// exchangesByIdOption is an interface that extends internal.Option to provide options specific to the derivatives exchanges by ID endpoint.
type exchangesByIdOption interface {
	internal.Option

	isExchangesByIdOption()
}

// WithIncludeTickersExchangesById returns an exchangesByIdOption that controls inclusion of tickers data.
// The includeTickers parameter specifies whether to include tickers data in the response.
func WithIncludeTickersExchangesById(includeTickers string) exchangesByIdOption {
	return includeTickersExchangesByIdOption{includeTickers}
}

// DerivativesExchangesId allows you to query the derivatives exchange’s related data (id, name, open interest, .... etc) based on the exchanges’ id.
//
//	👍 Tips
//
//	    You may change the value to either all to include all the tickers or unexpired to include unexpired tickers in the responses for the include_tickers params. You may leave it blank to omit the tickers data in response
//
//	📘Notes
//
//	    Cache / Update Frequency: every 30 seconds for all the API plans
func (c *DerivativesClient) DerivativesExchangesId(ctx context.Context, id string, options ...exchangesByIdOption) (*types.ExchangeId, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/exchanges/%s?%s", c.derivativesUrl(), id, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.ExchangeId
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type includeTickersExchangesByIdOption struct{ includeTickers string }

func (o includeTickersExchangesByIdOption) Apply(v *url.Values) {
	v.Set("include_tickers", o.includeTickers)
}
func (o includeTickersExchangesByIdOption) isExchangesByIdOption() {}
