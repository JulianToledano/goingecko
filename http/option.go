package http

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type (
	// ClientLike is a constraint interface that defines the common behavior
	// that both Client and RateLimitedClient should have
	ClientLike interface {
		*Client | *RateLimitedClient
	}

	// option is a generic function type that configures a client instance
	option[T ClientLike] func(T) T
)

// WithHttpClient returns an option that sets the HTTP client used by the client
func WithHttpClient[T ClientLike](client *http.Client) option[T] {
	return func(c T) T {
		switch v := any(c).(type) {
		case *Client:
			v.httpClient = client
		case *RateLimitedClient:
			v.httpClient = client
		}
		return c
	}
}

// WithApiHeaderFn returns an option that sets a function to add API headers to requests
func WithApiHeaderFn[T ClientLike](fn ApiHeaderFn) option[T] {
	return func(c T) T {
		switch v := any(c).(type) {
		case *Client:
			v.apiHeaderSetter = fn
		case *RateLimitedClient:
			v.apiHeaderSetter = fn
		}
		return c
	}
}

// WithRateLimit returns an option that sets a custom rate limiter for RateLimitedClient
// This option only works with RateLimitedClient - it's a no-op for regular Client
func WithRateLimit[T ClientLike](requestsPerMinute int) option[T] {
	return func(c T) T {
		switch v := any(c).(type) {
		case *RateLimitedClient:
			// Convert requests per minute to requests per second for the rate limiter
			rps := rate.Limit(float64(requestsPerMinute) / 60.0)
			v.rateLimiter = rate.NewLimiter(rps, 1)
		case *Client:
			// No-op for regular Client
		}
		return c
	}
}

// WithRateLimiter returns an option that sets a custom rate limiter instance for RateLimitedClient
// This option only works with RateLimitedClient - it's a no-op for regular Client
func WithRateLimiter[T ClientLike](limiter *rate.Limiter) option[T] {
	return func(c T) T {
		switch v := any(c).(type) {
		case *RateLimitedClient:
			v.rateLimiter = limiter
		}
		return c
	}
}

// WithRetryPolicy returns an option that sets a custom retry policy for RateLimitedClient.
// It configures how many times to retry failed requests and the base delay between retries.
// This option only works with RateLimitedClient - it's a no-op for regular Client.
//
// The retry policy uses exponential backoff, where each retry waits longer than the previous one.
// The delay for retry n is calculated as: baseDelay * 2^n
//
// For example, with baseDelay=2s:
// - First retry: 2s
// - Second retry: 4s
// - Third retry: 8s
// - Fourth retry: 16s
// - Fifth retry: 32s
func WithRetryPolicy[T ClientLike](maxRetries int, baseDelay time.Duration) option[T] {
	return func(c T) T {
		switch v := any(c).(type) {
		case *RateLimitedClient:
			v.retryPolicy = RetryPolicy{
				maxRetries: maxRetries,
				baseDelay:  baseDelay,
				withRetry:  true,
			}
		}
		return c
	}
}
