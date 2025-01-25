package git

import (
	"net/url"
	"os"
)

// NewEndpoint for GitHub.
func NewEndpoint() (*Endpoint, error) {
	endpoint := &Endpoint{}
	api := os.Getenv("GITHUB_API_URL")

	if api == "" {
		return endpoint, nil
	}

	u, err := url.Parse(api)
	if err != nil {
		return endpoint, err
	}

	endpoint.URL = u

	return endpoint, nil
}

// Endpoint for GitHub.
type Endpoint struct {
	*url.URL
}

// IsSet for GitHub.
func (e *Endpoint) IsSet() bool {
	return e.URL != nil
}
