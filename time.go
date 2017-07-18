package main

import (
	"time"

	. "github.com/reujab/bronze/types"
)

func timeSegment(segment *Segment) {
	segment.Value = time.Now().Format(time.Kitchen)
}
