package main

import (
	"os"

	. "github.com/reujab/bronze/types"
)

func envSegment(segment *Segment, args []string) {
	if len(args) != 1 {
		dief("env: expected exactly one argument")
	}

	segment.Value = os.Getenv(args[0])
}
