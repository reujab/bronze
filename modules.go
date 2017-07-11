package main

import (
	"fmt"
	"os"
)

func handleModule(module string, segment *segment) {
	switch module {
	case "os":
		osSegment(segment)
	case "status":
		statusSegment(segment)
	case "go":
		goSegment(segment)
	case "node":
		nodeSegment(segment)
	case "dir":
		dirSegment(segment)
	case "shortdir":
		shortDirSegment(segment)
	case "git":
		gitSegment(segment)
	case "cmdtime":
		cmdTimeSegment(segment)
	default:
		fmt.Fprintf(os.Stderr, "bronze: Invalid module: %q.\n", module)
		os.Exit(1)
	}
}
