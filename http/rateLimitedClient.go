package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// HttpRateLimitedlient is an interface that extends HttpClient to provide rate limiting functionality.
// It adds the ability to query the current rate limit configuration and state.
type HttpRateLimitedlient interface {
	HttpClient
	// GetRateLimitInfo returns the current rate limit configuration and state.
	// It returns:
	//   - limit: The maximum number of requests allowed per second
	//   - burst: The maximum number of requests that can be made in a single burst
	//   - tokens: The current number of available tokens
	GetRateLimitInfo() (limit rate.Limit, burst int, tokens int)
}

var (
	_ HttpClient           = &RateLimitedClient{}
	_ HttpRateLimitedlient = &RateLimitedClient{}
)

// RetryPolicy defines the configuration for retrying failed requests.
// It specifies how many times to retry and the base delay between retries.
type RetryPolicy struct {
	// maxRetries is the maximum number of times to retry a failed request
	maxRetries int
	// baseDelay is the initial delay between retries, which will be exponentially increased
	baseDelay time.Duration
	// withRetry is a boolean that indicates if the retry policy should be used
	withRetry bool
}

// defaultRetryPolicy defines the default configuration for retrying failed requests.
// It allows up to 5 retries with an initial delay of 2 seconds between attempts.
// The delay will be exponentially increased for each subsequent retry.
var defaultRetryPolicy = RetryPolicy{
	maxRetries: 5,
	baseDelay:  2 * time.Second,
	withRetry:  true,
}

// RateLimitedClient extends Client to add rate limiting functionality.
// It uses a token bucket rate limiter to control request frequency.
type RateLimitedClient struct {
	Client
	rateLimiter *rate.Limiter
	retryPolicy RetryPolicy
}

// NewRateLimitedClient creates a new RateLimitedClient with the given options
// By default, it applies CoinGecko's free tier rate limit of 30 requests per minute
func NewRateLimitedClient(opts ...option[*RateLimitedClient]) *RateLimitedClient {
	// Default rate limiter for CoinGecko free tier: 30 requests per minute
	// Using conservative 10 requests per minute to avoid 429s
	defaultRateLimit := rate.NewLimiter(rate.Limit(10.0/60.0), 1)

	c := &RateLimitedClient{
		rateLimiter: defaultRateLimit,
		retryPolicy: defaultRetryPolicy,
	}

	for _, opt := range opts {
		c = opt(c)
	}

	// Set default HTTP client if none provided
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return c
}

// MakeReq makes a GET request to the given URL with the configured client
// It respects the configured rate limit before making the request
// It returns the response body as bytes or an error if the request fails
func (c *RateLimitedClient) MakeReq(ctx context.Context, url string) ([]byte, error) {
	for attempt := 0; attempt <= c.retryPolicy.maxRetries; attempt++ {
		// Wait for rate limiter permission before making the request
		if err := c.rateLimiter.Wait(ctx); err != nil {
			return nil, fmt.Errorf("rate limiter wait failed: %w", err)
		}

		resp, err := c.Client.MakeReq(ctx, url)
		if err != nil {
			if !c.retryPolicy.withRetry {
				return resp, err
			}
			// Check if this is a 429 error (rate limit exceeded)
			if apiErr, ok := err.(*APIError); ok && apiErr.StatusCode == 429 {
				if attempt < c.retryPolicy.maxRetries {
					// Calculate exponential backoff delay
					delay := c.retryPolicy.baseDelay * time.Duration(1<<attempt) // 2s, 4s, 8s, 16s, 32s

					select {
					case <-time.After(delay):
						continue // retry
					case <-ctx.Done():
						return nil, fmt.Errorf("context cancelled during retry wait: %w", ctx.Err())
					}
				}
				// If we've exhausted retries, return the 429 error
				return nil, fmt.Errorf("max retries exceeded for rate limiting: %w", err)
			}
			// For non-429 errors, return immediately
			return nil, err
		}
		// Success, return the response
		return resp, nil
	}

	// This should never be reached, but just in case
	return nil, fmt.Errorf("unexpected error: max retries exceeded")
}

// GetRateLimitInfo returns information about the current rate limiter state
func (c *RateLimitedClient) GetRateLimitInfo() (limit rate.Limit, burst int, tokens int) {
	if c.rateLimiter == nil {
		return 0, 0, 0
	}

	limit = c.rateLimiter.Limit()
	burst = c.rateLimiter.Burst()
	tokens = int(c.rateLimiter.TokensAt(time.Now()))

	return limit, burst, tokens
}
