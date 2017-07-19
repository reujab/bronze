package main

import (
	"fmt"
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

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func dief(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "bronze: "+format, args...)
	os.Exit(1)
}

func dieIf(err error, format string, args ...interface{}) {
	if err != nil {
		dief(format, args...)
	}
}
