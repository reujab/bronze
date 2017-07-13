package main

import (
	"fmt"
	"os"
	"time"
)

func cmdTimeSegment(segment *segment) {
	duration, err := time.ParseDuration(os.Getenv("cmdtime"))
	if err != nil {
		if shell == "bash" {
			fmt.Fprintln(os.Stderr, "bronze: The 'cmdtime' module is supported in bash.")
			os.Exit(1)
		}
		panic(err)
	}

	threshold, err := time.ParseDuration(os.Getenv("BRONZE_CMDTIME_THRESHOLD"))
	if err != nil || duration >= threshold {
		segment.value = duration.String()
	}
}
