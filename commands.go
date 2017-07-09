package main

import (
	"fmt"
	"os"
)

func handleCommand(command string, segment *segment) {
	switch command {
	case "dir":
		dirSegment(segment)
	default:
		fmt.Fprintf(os.Stderr, "bronze: Invalid command: %q.\n", command)
		os.Exit(1)
	}
}
