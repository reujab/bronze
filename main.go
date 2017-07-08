package main

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var shell = filepath.Base(os.Getenv("SHELL"))

func main() {
	app := cli.NewApp()
	app.Usage = "a cross-shell customizable powerline-like prompt with icons"
	app.Commands = []cli.Command{
		cmdPrint,
	}
	app.Run(os.Args)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
