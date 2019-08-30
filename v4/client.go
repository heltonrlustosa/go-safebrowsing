package v4

import (
	"net/http"

	"golang.org/x/time/rate"
)

// Client may be used to do request to the Google Safe Browsing APIs
type Client struct {
	httpClient        *http.Client
	apiKey            string
	baseURL           string
	clientID          string
	signature         []byte
	requestsPerSecond int
	rateLimiter       *rate.Limiter
	channel           string
}

// ClientOption is the type of constructor options for NewClient(...).
type ClientOption func(*Client) error
