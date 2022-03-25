package main

import (
	"os"
	"time"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

// nolint:gomnd
func main() {
	command, err := scmd.New(15*time.Second, cmd.ServeOptions, cmd.WorkerOptions)
	if err != nil {
		os.Exit(1)
	}

	if err := command.Execute(); err != nil {
		os.Exit(2)
	}
}
