package main

//go:generate go get github.com/UnnoTed/fileb0x
//go:generate fileb0x b0x.yaml

import (
	"os"

	"github.com/reujab/bronze/static"
	"github.com/urfave/cli"
)

var cmdInit = cli.Command{
	Name:        "init",
	Usage:       "Initializes the shell for use of bronze",
	Description: `Put 'eval "$(bronze init)"' in your ~/.*shrc`,
	Action: func(ctx *cli.Context) error {
		script, err := static.ReadFile("init." + shell)
		if err == nil {
			os.Stdout.Write(script)
		} else {
			dief("unrecognized shell")
		}
		return nil
	},
}
