package v4

// ClientInfo is a required request client information
type ClientInfo struct {
	ClientID      string `json:"clientId"`
	ClientVersion string `json:"clientVersion"`
}

// ThreatInfo defines thread information
// TODO: describe all fields
type ThreatInfo struct {
	ThreatTypes      []string `json:"threatTypes"`
	PlatformTypes    []string `json:"platformTypes"`
	ThreatEntryTypes []string `json:"threatEntryTypes"`
	ThreatEntries    []string `json:"threatEntries"`
}

// Entry defines thread metadata
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ThreatMatch is a default structure response when successful
type ThreatMatch struct {
	// The threat type matching this threat
	ThreatType string `json:"threatType"`

	// The platform type matching this threat
	PlatformType string `json:"platformType"`

	// The threat entry type matching this threat
	ThreatEntryType string `json:"threatEntryType"`

	// The threat matching this threat
	Threat map[string]interface{} `json:"threat"`

	// Optional metadata associated with this threat
	ThreatEntryMetadata map[string][]Entry `json:"threatEntryMetadata"`

	// The cache lifetime for the returned match Clients must not cache this
	// response for more than this duration to avoid false positives
	CacheDuration string `json:"cacheDuration"`
}
