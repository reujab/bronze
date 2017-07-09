package main

import (
	"fmt"
	"os"
)

func handleCommand(command string, segment *segment) {
	switch command {
	case "status":
		statusSegment(segment)
	case "dir":
		dirSegment(segment)
	case "git":
		gitSegment(segment)
	case "cmdtime":
		cmdTimeSegment(segment)
	default:
		fmt.Fprintf(os.Stderr, "bronze: Invalid command: %q.\n", command)
		os.Exit(1)
	}
}
