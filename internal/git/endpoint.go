package git

import (
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/strings"
)

// NewEndpoint for GitHub.
func NewEndpoint() Endpoint {
	return Endpoint(os.GetVariable("GITHUB_API_URL"))
}

// Endpoint for GitHub.
type Endpoint string

// IsSet for GitHub.
func (e Endpoint) IsSet() bool {
	return !strings.IsEmpty(e.String())
}

// String conforms to fmt.Stringer.
func (e Endpoint) String() string {
	return string(e)
}
