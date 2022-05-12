//go:build features

package main

import (
	"testing"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

func TestFeatures(t *testing.T) {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)
	command.AddClient(cmd.ClientOptions)
	command.AddVersion(cmd.Version)

	if err := command.Run(); err != nil {
		t.Fatal(err.Error())
	}
}
