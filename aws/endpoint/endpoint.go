package endpoint

import (
	"os"
)

// NewEndpoint for AWS.
func NewEndpoint() Endpoint {
	return Endpoint(os.Getenv("AWS_URL"))
}

// Endpoint for AWS.
type Endpoint string

// IsSet for AWS.
func (e Endpoint) IsSet() bool {
	return e != ""
}
