package git

import (
	"os"
)

// NewEndpoint for GitHub.
func NewEndpoint() Endpoint {
	return Endpoint(os.Getenv("GITHUB_API_URL"))
}

// Endpoint for GitHub.
type Endpoint string

// IsSet for GitHub.
func (e Endpoint) IsSet() bool {
	return e != ""
}
