package main

import (
	"os"
	"path"
	. "github.com/reujab/bronze/types"
)

func virtualEnvSegment(segment *Segment) {
	content := os.Getenv("VIRTUAL_ENV")
	if (content != "") {
		segment.Value = path.Base(content)
	}
}
