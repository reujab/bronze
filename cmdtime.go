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

	segment.value = duration.String()
}
