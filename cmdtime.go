package main

import (
	"fmt"
	"os"
	"time"
)

func cmdTimeSegment(segment *segment) {
	duration, err := time.ParseDuration(os.Getenv("CMDTIME"))
	if err != nil {
		if shell == "bash" {
			fmt.Fprintln(os.Stderr, "bronze: The 'cmdtime' command is supported in bash.")
		}
		panic(err)
	}

	segment.value = duration.String()
}
