package cmd

import (
	"github.com/alexfalkowski/go-service/env"
)

// Version of the app.
var Version = "development"

// NewVersion of the app.
func NewVersion() env.Version {
	return env.Version(Version)
}
