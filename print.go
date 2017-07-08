package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/urfave/cli"
)

// powerline triangle
const separator = "\ue0b0"

var cmdPrint = cli.Command{
	Name:  "print",
	Usage: "Prints the prompt with the specified syntax",
	// TODO: improve description
	Description: "Example: bronze print blue:black:dir",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			cli.ShowCommandHelpAndExit(ctx, "print", 1)
		}

		cmdPrintAction(ctx.Args())
		return nil
	},
}

type segment struct {
	visible    bool
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
		command := fields[2]

		segment := &segment{
			background: fields[0],
			foreground: fields[1],
		}
		segments = append(segments, segment)

		switch command {
		case "dir":
			go func() {
				dirSegment(segment)
				waitgroup.Done()
			}()
		default:
			fmt.Fprintf(os.Stderr, "bronze: Invalid command: %q.\n", command)
			os.Exit(1)
		}
	}

	// wait for all the async segments
	waitgroup.Wait()

	// print the prompt
	first := true
	lastBackground := "white"
	for _, segment := range segments {
		if !segment.visible {
			continue
		}

		// if this isn't the first segment, before printing the next segment, separate them
		if !first {
			// use the last background as the current foreground
			printSegment(segment.background, lastBackground, separator)
		}
		first = false

		printSegment(segment.background, segment.foreground, " "+segment.value+" ")
		lastBackground = segment.background
	}
	// print final separator
	printSegment("reset", lastBackground, separator)
}

func printSegment(background, foreground, value string) {
	fmt.Print(escapeBackground(background) + escapeForeground(foreground) + value)
}
