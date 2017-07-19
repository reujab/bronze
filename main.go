package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var shell = os.Getenv("BRONZE_SHELL")

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
