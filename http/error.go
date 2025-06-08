package http

import "fmt"

// APIError is an error returned by the client when the API returns a non-200 status code
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api returned status code %d", e.StatusCode)
}
