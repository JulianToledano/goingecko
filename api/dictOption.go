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

func (o dictOption) IsCoinsIdOption() {}
