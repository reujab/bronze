package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	. "github.com/reujab/bronze/types"
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

func cmdPrintAction(args []string) {
	var segments []*Segment
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(len(args))

	for _, arg := range args {
		// validate argument
		fields := strings.Split(arg, ":")
		if len(fields) < 3 {
			fmt.Fprintf(os.Stderr, "bronze: Invalid argument: %q. At least three fields expected.\n", arg)
			os.Exit(1)
		}

		segment := &Segment{
			Background: fields[1],
			Foreground: fields[2],
		}
		segments = append(segments, segment)

		go func() {
			handleModule(fields[0], segment, fields[3:])
			waitgroup.Done()
		}()
	}

	// wait for all the async segments
	waitgroup.Wait()

	// print the prompt
	first := true
	lastSegment := &Segment{
		Background: "black",
		Foreground: "white",
	}
	for _, segment := range segments {
		if segment.Value == "" {
			continue
		}

		// if this isn't the first segment, before printing the next segment, separate them
		if !first {
			if segment.Background == lastSegment.Background {
				printSegment(segment.Background, lastSegment.Foreground, thinSeparator)
			} else {
				// use the last background as the current foreground
				printSegment(segment.Background, lastSegment.Background, separator)
			}
		}
		first = false

		printSegment(segment.Background, segment.Foreground, " "+segment.Value+" ")
		lastSegment = segment
	}
	// print final separator
	printSegment("none", lastSegment.Background, separator)
	resetColors()
}

func printSegment(background, foreground, value string) {
	fmt.Print(escapeBackground(background) + escapeForeground(foreground) + value)
}
