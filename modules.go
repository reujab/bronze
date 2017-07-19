package main

import (
	"fmt"
	"os"

	. "github.com/reujab/bronze/types"
)

func handleModule(module string, segment *Segment, args []string) {
	switch module {
	case "os":
		osSegment(segment)
	case "status":
		statusSegment(segment)
	case "packages":
		packagesSegment(segment)
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
	case "plugin":
		pluginSegment(segment, args)
	default:
		fmt.Fprintf(os.Stderr, "bronze: Invalid module: %q.\n", module)
		os.Exit(1)
	}
}
