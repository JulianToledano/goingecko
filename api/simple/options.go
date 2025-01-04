package simple

import (
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// priceOption is an interface that extends internal.Option with a marker method
// to ensure type safety for price-specific options
type priceOption interface {
	internal.Option
	isPriceOption()
}

// WithIncludeMarketCapOption returns a priceOption that configures whether to include
// market cap data in the price response
func WithIncludeMarketCapOption(includeMarketCap bool) priceOption {
	return includeMarketCapOption{includeMarketCap}
}

// WithIncludeDayVolumeOption returns a priceOption that configures whether to include
// 24-hour volume data in the price response
func WithIncludeDayVolumeOption(includeDayVolume bool) priceOption {
	return includeDayVolumeOption{includeDayVolume}
}

// WithIncludeDayChangeOption returns a priceOption that configures whether to include
// 24-hour price change data in the price response
func WithIncludeDayChangeOption(includeDayChange bool) priceOption {
	return includeDayChangeOption{includeDayChange}
}

// WithIncludeLastTimeUpdatedAtOption returns a priceOption that configures whether to include
// the last update timestamp in the price response
func WithIncludeLastTimeUpdatedAtOption(includeLastTimeUpdatedAt bool) priceOption {
	return includeLastTimeUpdatedOption{includeLastTimeUpdatedAt}
}

// WithPrecisionOption returns a priceOption that configures the decimal precision
// of price values in the response
func WithPrecisionOption(precision string) priceOption {
	return precisionOption{precision}
}

type (
	includeMarketCapOption       struct{ include bool }
	includeDayVolumeOption       struct{ include bool }
	includeDayChangeOption       struct{ include bool }
	includeLastTimeUpdatedOption struct{ include bool }
	precisionOption              struct{ precision string }
)

func (o includeMarketCapOption) Apply(v *url.Values) {
	v.Set("include_market_cap", strconv.FormatBool(o.include))
}

func (o includeDayVolumeOption) Apply(v *url.Values) {
	v.Set("include_24hr_vol", strconv.FormatBool(o.include))
}

func (o includeDayChangeOption) Apply(v *url.Values) {
	v.Set("include_24hr_change", strconv.FormatBool(o.include))
}

func (o includeLastTimeUpdatedOption) Apply(v *url.Values) {
	v.Set("include_last_updated_at", strconv.FormatBool(o.include))
}

func (o precisionOption) Apply(v *url.Values) {
	v.Set("precision", o.precision)
}

func (o includeMarketCapOption) isPriceOption()       {}
func (o includeDayVolumeOption) isPriceOption()       {}
func (o includeDayChangeOption) isPriceOption()       {}
func (o includeLastTimeUpdatedOption) isPriceOption() {}
func (o precisionOption) isPriceOption()              {}
