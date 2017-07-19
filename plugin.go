package main

import (
	"path/filepath"
	"plugin"

	. "github.com/reujab/bronze/types"
)

func pluginSegment(segment *Segment, args []string) {
	if len(args) == 0 {
		dief("plugin: expected at least one argument")
	}

	plug, err := plugin.Open(args[0])
	dieIf(err, "plugin: failed to open plugin: %q", args[0])

	handler, err := plug.Lookup("Main")
	dieIf(err, "plugin: failed to lookup Main symbol in %q", filepath.Base(args[0]))
	handler.(func(*Segment, []string))(segment, args[1:])
}
