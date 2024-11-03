package goingecko

import (
	"net/url"
)

// Option is a generic option type
type Option interface {
	apply(*url.Values)
}
