package main

import (
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
	case "rss":
		rssSegment(segment)
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
	case "env":
		envSegment(segment, args)
	case "plugin":
		pluginSegment(segment, args)
	default:
		dief("invalid module: %q", module)
	}
}
