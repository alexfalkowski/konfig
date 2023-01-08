package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/marshaller"
)

// OutputFlag for client cmd.
var OutputFlag string

// OutputConfig for cmd.
type OutputConfig struct {
	*cmd.Config
}

// NewOutputConfig for cmd.
func NewOutputConfig(factory *marshaller.Factory) (*OutputConfig, error) {
	c, err := cmd.NewConfig(OutputFlag, factory)
	if err != nil {
		return nil, err
	}

	return &OutputConfig{Config: c}, nil
}
