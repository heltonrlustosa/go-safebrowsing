package v4

import (
	"context"
	"os"
)

var apikey string

func init() {
	key := os.Getenv("GOOGLE_API_KEY")
	if key != "" {
		apikey = key
	}
}

const (
	// Lookup API
	findThreadMatches = "/v4/threatMatches:find"

	// Update API
	findHashPath    = "/v4/fullHashes:find"
	fetchUpdatePath = "/v4/threatListUpdates:fetch"
)

type apiV4 interface {
	FindThreadMatches(ctx context.Context, req FindThreatMatchesRequest) (*FindThreatMatchesResponse, error)

	// TODO:
	// FindHashPath(ctx context.Context, req FindHashPathRequest)
	// FetchUpdatePath(ctx context.Context, req FetchUpdatePathRequest)
}
