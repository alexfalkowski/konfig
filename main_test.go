//go:build features
// +build features

package main

import (
	"testing"
	"time"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

func TestFeatures(t *testing.T) {
	command, err := scmd.New(15*time.Second, cmd.ServerOptions, cmd.WorkerOptions)
	if err != nil {
		t.Fatal(err.Error())
	}

	if err := command.Execute(); err != nil {
		t.Fatal(err.Error())
	}
}
