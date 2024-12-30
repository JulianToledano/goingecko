package contract

import (
	"net/url"

	"github.com/JulianToledano/goingecko/v3/api/internal"
)

// contractOption is an interface that extends internal.Option to provide
// contract-specific market chart options
type contractOption interface {
	internal.Option

	isMarketChartOption()
}

// WithIntervalOption returns a contractOption that sets the interval parameter
// for market chart data requests
func WithIntervalOption(interval string) contractOption {
	return intervalContractOption{interval}
}

// WithPrecisionOption returns a contractOption that sets the precision parameter
// for market chart data requests
func WithPrecisionOption(precision string) contractOption {
	return precisionContractOption{precision}
}

type intervalContractOption struct{ interval string }
type precisionContractOption struct{ precision string }

func (o intervalContractOption) Apply(v *url.Values) {
	v.Set("interval", o.interval)
}
func (o precisionContractOption) Apply(v *url.Values) {
	v.Set("precision", o.precision)
}

func (o intervalContractOption) isMarketChartOption()  {}
func (o precisionContractOption) isMarketChartOption() {}
