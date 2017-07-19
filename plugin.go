package main

import (
	"fmt"
	"os"
	"plugin"

	. "github.com/reujab/bronze/types"
)

func pluginSegment(segment *Segment, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "bronze: plugin: Expected at least one argument.")
		os.Exit(1)
	}

	plug, err := plugin.Open(args[0])
	die(err)

	handler, err := plug.Lookup("Main")
	die(err)
	handler.(func(*Segment, []string))(segment, args[1:])
}
