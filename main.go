package main

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

// the shell environment variable is not always the current shell
// you should explicitly set the variable in your *shrc
var shell = filepath.Base(os.Getenv("SHELL"))

func main() {
	app := cli.NewApp()
	app.Usage = "a cross-shell customizable powerline-like prompt with icons"
	app.Commands = []cli.Command{
		cmdInit,
		cmdPrint,
	}
	app.Run(os.Args)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
