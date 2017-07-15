package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/urfave/cli"
)

var cmdPrint = cli.Command{
	Name:  "print",
	Usage: "Prints the prompt with the specified syntax",
	// TODO: improve description
	Description: "Example: bronze print dir:blue:black",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			cli.ShowCommandHelpAndExit(ctx, "print", 1)
		}

		cmdPrintAction(ctx.Args())
		return nil
	},
}

type segment struct {
	background string
	foreground string
	value      string
}

func cmdPrintAction(args []string) {
	var segments []*segment
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(len(args))

	for _, arg := range args {
		// validate argument
		fields := strings.Split(arg, ":")
		if len(fields) != 3 {
			fmt.Fprintf(os.Stderr, "bronze: Invalid argument: %q. Exactly three fields expected.\n", arg)
			os.Exit(1)
		}

		segment := &segment{
			background: fields[1],
			foreground: fields[2],
		}
		segments = append(segments, segment)

		go func() {
			handleModule(fields[0], segment)
			waitgroup.Done()
		}()
	}

	// wait for all the async segments
	waitgroup.Wait()

	// print the prompt
	first := true
	lastSegment := &segment{
		background: "black",
		foreground: "white",
	}
	for _, segment := range segments {
		if segment.value == "" {
			continue
		}

		// if this isn't the first segment, before printing the next segment, separate them
		if !first {
			if segment.background == lastSegment.background {
				printSegment(segment.background, lastSegment.foreground, "\ue0b1")
			} else {
				// use the last background as the current foreground
				printSegment(segment.background, lastSegment.background, separator)
			}
		}
		first = false

		printSegment(segment.background, segment.foreground, " "+segment.value+" ")
		lastSegment = segment
	}
	// print final separator
	printSegment("none", lastSegment.background, separator)
	resetColors()
}

func printSegment(background, foreground, value string) {
	fmt.Print(escapeBackground(background) + escapeForeground(foreground) + value)
}
