package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/konfig/internal/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(env.NewName(), env.NewVersion())

	cmd.RegisterServer(command)

	return command
}
