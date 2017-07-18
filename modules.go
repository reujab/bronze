package main

import (
	"fmt"
	"os"

	. "github.com/reujab/bronze/types"
)

func handleModule(module string, segment *Segment) {
	switch module {
	case "os":
		osSegment(segment)
	case "status":
		statusSegment(segment)
	case "pacaur":
		pacaurSegment(segment)
	case "user":
		userSegment(segment)
	case "dir":
		dirSegment(segment)
	case "git":
		gitSegment(segment)
	case "cmdtime":
		cmdTimeSegment(segment)
	case "time":
		timeSegment(segment)
	default:
		fmt.Fprintf(os.Stderr, "bronze: Invalid module: %q.\n", module)
		os.Exit(1)
	}
}
