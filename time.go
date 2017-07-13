package main

import "time"

func timeSegment(segment *segment) {
	segment.value = time.Now().Format(time.Kitchen)
}
