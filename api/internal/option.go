package internal

import "net/url"

// Option is a generic option type
type Option interface {
	Apply(*url.Values)
}
