package main

import (
	"os"
	"path"

	. "github.com/reujab/bronze/types"
)

func virtualenvSegment(segment *Segment) {
	segment.Value = path.Base(os.Getenv("VIRTUAL_ENV"))
}
