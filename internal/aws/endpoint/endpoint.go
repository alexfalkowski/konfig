package endpoint

import (
	"github.com/alexfalkowski/go-service/os"
)

// NewEndpoint for AWS.
func NewEndpoint() Endpoint {
	return Endpoint(os.GetVariable("AWS_URL"))
}

// Endpoint for AWS.
type Endpoint string

// IsSet for AWS.
func (e Endpoint) IsSet() bool {
	return e != ""
}
