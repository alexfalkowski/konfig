package endpoint

import (
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/strings"
)

// NewEndpoint for AWS.
func NewEndpoint() Endpoint {
	return Endpoint(os.GetVariable("AWS_URL"))
}

// Endpoint for AWS.
type Endpoint string

// IsSet for AWS.
func (e Endpoint) IsSet() bool {
	return !strings.IsEmpty(e.String())
}

// String conforms to fmt.Stringer.
func (e Endpoint) String() string {
	return string(e)
}
