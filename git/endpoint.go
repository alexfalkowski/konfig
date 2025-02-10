package git

import (
	"github.com/alexfalkowski/go-service/os"
)

// NewEndpoint for GitHub.
func NewEndpoint() Endpoint {
	return Endpoint(os.GetVariable("GITHUB_API_URL"))
}

// Endpoint for GitHub.
type Endpoint string

// IsSet for GitHub.
func (e Endpoint) IsSet() bool {
	return e != ""
}
