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

var colors = map[string]string{
	"fg-black":   "30",
	"fg-red":     "31",
	"fg-green":   "32",
	"fg-yellow":  "33",
	"fg-blue":    "34",
	"fg-magenta": "35",
	"fg-cyan":    "36",
	"fg-white":   "37",
	"bg-black":   "40",
	"bg-red":     "41",
	"bg-green":   "42",
	"bg-yellow":  "43",
	"bg-blue":    "44",
	"bg-magenta": "45",
	"bg-cyan":    "46",
	"bg-white":   "47",
}

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
	background string
	foreground string
	value      string
}

func cmdPrintAction(args []string) {
	var segments []segment
	waitgroup := new(sync.WaitGroup)
	waitgroup.Add(len(args))

	for _, arg := range args {
		// validate argument
		fields := strings.Split(arg, ":")
		if len(fields) != 3 {
			fmt.Fprintf(os.Stderr, "bronze: Invalid argument: %q. Exactly three fields expected.\n", arg)
			os.Exit(1)
		}
		_, ok := colors["bg-"+fields[0]]
		if !ok {
			fmt.Fprintf(os.Stderr, "bronze: Invalid background: %q.\n", fields[0])
			os.Exit(1)
		}
		_, ok = colors["fg-"+fields[1]]
		if !ok {
			fmt.Fprintf(os.Stderr, "bronze: Invalid foreground: %q.\n", fields[1])
			os.Exit(1)
		}
		command := fields[2]

		segments = append(segments, segment{
			background: fields[0],
			foreground: fields[1],
		})

		// TODO: make async
		// FIXME: temporary code
		switch command {
		case "dir":
			dir, err := os.Getwd()
			check(err)
			segments[len(segments)-1].value = dir
			waitgroup.Done()
		default:
			fmt.Fprintf(os.Stderr, "bronze: Invalid command: %q.\n", command)
			os.Exit(1)
		}
	}

	// wait for all the async segments
	waitgroup.Wait()

	// print the prompt
	for i, segment := range segments {
		// if this isn't the first segment, before printing the next segment, separate them
		if i != 0 {
			// use the last background as the current foreground
			printSegment(segment.background, segments[i-1].background, separator)
		}

		printSegment(segment.background, segment.foreground, " "+segment.value+" ")
	}
	// print final separator
	printSegment("black", segments[len(segments)-1].background, separator)
}

func printSegment(background, foreground, value string) {
	fmt.Printf("\x1b[%sm\x1b[%sm%s", colors["bg-"+background], colors["fg-"+foreground], value)
}
