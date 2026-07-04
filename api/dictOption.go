package api

import (
	"fmt"
	"net/url"
)

type dictOption map[string]any

// WithURLParams returns an option that applies the provided key-value pairs as
// URL query parameters. It can be passed to any API method that accepts options.
func WithURLParams(params map[string]any) dictOption {
	return dictOption(params)
}

func (o dictOption) Apply(v *url.Values) {
	for k, value := range o {
		v.Set(k, fmt.Sprintf("%v", value))
	}
}

func (o dictOption) IsAssetPlatformsOption()           {}
func (o dictOption) IsCategoryOptions()                {}
func (o dictOption) IsCoinsIdOption()                  {}
func (o dictOption) IsCoinsIdTickersOption()           {}
func (o dictOption) IsEntitiesListOption()             {}
func (o dictOption) IsEntityOption()                   {}
func (o dictOption) IsExchangesByIdOption()            {}
func (o dictOption) IsExchangesOption()                {}
func (o dictOption) IsExchangesOptions()               {}
func (o dictOption) IsHoldingChartOption()             {}
func (o dictOption) IsIdCirculatingSupplyChartOption() {}
func (o dictOption) IsIdHistoryOption()                {}
func (o dictOption) IsIdMarketChartOption()            {}
func (o dictOption) IsIdMarketChartRangeOption()       {}
func (o dictOption) IsIdTotalSupplyChartOption()       {}
func (o dictOption) IsListOption()                     {}
func (o dictOption) IsMarketCapChartOption()           {}
func (o dictOption) IsMarketChartOption()              {}
func (o dictOption) IsMarketsOption()                  {}
func (o dictOption) IsOhlcOption()                     {}
func (o dictOption) IsPriceOption()                    {}
func (o dictOption) IsPublicTreasuryCoinIdOption()     {}
func (o dictOption) IsTickersOption()                  {}
func (o dictOption) IsTopGainersLosersOption()         {}
func (o dictOption) IsTransactionHistoryOption()       {}
