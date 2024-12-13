package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/konfig/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:KONFIG_CONFIG_FILE")
	c.AddServer("server", "Start konfig server", cmd.ServerOptions...)

	return c
}
