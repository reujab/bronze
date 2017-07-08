package main

import "github.com/urfave/cli"

var cmdPrint = cli.Command{
	Name:        "print",
	Usage:       "Prints the prompt with the specified syntax",
	Description: "Example: bronze print blue:black:dir",
	Action: func(ctx *cli.Context) error {
		// TODO
		return nil
	},
}
