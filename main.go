package main

import (
	"os"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/client"
	"github.com/alexfalkowski/konfig/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New(cmd.Version)
	command.AddServer(cmd.ServerOptions...)

	c := command.AddClient(cmd.ClientOptions...)
	c.PersistentFlags().StringVarP(
		&client.OutputFlag,
		"output", "o", "env:APP_CONFIG_FILE", "output config location (format kind:location, default env:APP_CONFIG_FILE)",
	)

	return command
}
