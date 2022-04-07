//go:build features

package main

import (
	"testing"
	"time"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

func TestFeatures(t *testing.T) {
	command := scmd.New(15 * time.Second)

	command.AddServer(cmd.ServerOptions)
	command.AddWorker(cmd.WorkerOptions)
	command.AddClient(cmd.ClientOptions)

	if err := command.Run(); err != nil {
		t.Fatal(err.Error())
	}
}
