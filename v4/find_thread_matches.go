package v4

// FindThreatMatchesRequest is a request body to check entries against lists
type FindThreatMatchesRequest struct {
	// The client metadata
	ClientInfo ClientInfo `json:"clientInfo"`

	// The lists and entries to be checked for matches
	ThreatInfo ThreatInfo `json:"threatInfo"`
}

// FindThreatMatchesResponse is a response body to check entries
type FindThreatMatchesResponse struct {
	// The threat list matches
	Matches []ThreatMatch          `json:"matches"`
	Error   map[string]interface{} `json:"error"`
}
