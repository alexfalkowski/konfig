package main

import (
	"os"
	"time"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

// nolint:gomnd
func main() {
	command, err := scmd.New(15 * time.Second)
	if err != nil {
		os.Exit(1)
	}

	command.AddServer(cmd.ServerOptions)
	command.AddWorker(cmd.WorkerOptions)
	command.AddClient(cmd.ClientOptions)

	if err := command.Run(); err != nil {
		os.Exit(2)
	}
}
