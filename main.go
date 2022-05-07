package main

import (
	"os"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

func main() {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)
	command.AddWorker(cmd.WorkerOptions)
	command.AddClient(cmd.ClientOptions)
	command.AddVersion(cmd.Version)

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
