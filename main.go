package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "a cross-shell customizable powerline-like prompt with icons"
	app.Commands = []cli.Command{
		cmdPrint,
	}
	app.Run(os.Args)
}
