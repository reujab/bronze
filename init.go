package main

//go:generate go get github.com/jteeuwen/go-bindata/go-bindata
//go:generate go-bindata -nocompress init.bash init.zsh init.fish

import (
	"os"

	"github.com/urfave/cli"
)

var cmdInit = cli.Command{
	Name:        "init",
	Usage:       "Initializes the shell for use of bronze",
	Description: `Put 'eval "$(bronze init)"' in your ~/.*shrc`,
	Action: func(ctx *cli.Context) error {
		script, err := Asset("init." + shell)
		if err == nil {
			os.Stdout.Write(script)
		} else {
			dief("unrecognized shell")
		}
		return nil
	},
}
