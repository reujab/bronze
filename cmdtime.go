package main

import (
	"os"
	"time"

	. "github.com/reujab/bronze/types"
)

func cmdTimeSegment(segment *Segment) {
	duration, err := time.ParseDuration(os.Getenv("cmdtime"))
	if err != nil {
		if shell == "bash" {
			dief("cmdtime: bash is not supported")
		} else {
			dief("cmdtime: invalid $cmdtime: %q", os.Getenv("cmdtime"))
		}
	}

	threshold, err := time.ParseDuration(os.Getenv("BRONZE_CMDTIME_THRESHOLD"))
	if err != nil || duration >= threshold {
		segment.Value = duration.String()
	}
}
