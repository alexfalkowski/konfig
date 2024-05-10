package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/konfig/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	command := sc.New(cmd.Version)
	command.AddServer(cmd.ServerOptions...)

	c := command.AddClient(cmd.ClientOptions...)
	flags.StringVar(c, sc.OutputFlag,
		"output", "o", "env:APP_CONFIG_FILE", "output config location (format kind:location, default env:APP_CONFIG_FILE)")

	return command
}
