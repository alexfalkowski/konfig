package main

import (
	"os"

	scmd "github.com/alexfalkowski/go-service/cmd"
	ccmd "github.com/alexfalkowski/konfig/client/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)

	c := command.AddClient(cmd.ClientOptions)
	c.PersistentFlags().StringVar(
		&ccmd.OutputFlag,
		"output", "env:APP_CONFIG_FILE", "output config location (format kind:location, default env:APP_CONFIG_FILE)",
	)

	command.AddVersion(cmd.Version)

	return command
}
