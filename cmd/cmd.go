package cmd

import (
	"github.com/alexfalkowski/go-service/version"
)

// Version of the app.
var Version = "development"

// NewVersion of the app.
func NewVersion() version.Version {
	return version.Version(Version)
}
